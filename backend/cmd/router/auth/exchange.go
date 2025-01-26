package auth_router

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleExchange(w http.ResponseWriter, req *http.Request) {
	dto := &dtos.Exchange{}
	err := json.NewDecoder(req.Body).Decode(dto)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = dto.Validate(r.v)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	es, exists := r.auth.MagicLink.GetSession(dto.Email)
	if !exists {
		logger.WarnC("email sessions doesn't exist")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if es.ActivateCode != dto.Code {
		logger.WarnC("provided a wrong code ", es.UserEmail)
		logger.Debug(es.ActivateCode, dto.Code)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if es.ExchangeCounter > 5 || es.IsBlocked {
		es.IsBlocked = true
		logger.Warn("exchange counter block", es.UserEmail)
		http.Error(w, "Your account is blocked due the security reasones.", http.StatusBadRequest)
		return
	}

	es.ExchangeCounter += 1
	u, err := r.d.GetUserByEmail(req.Context(), sql.NullString{
		String: es.UserEmail,
	})
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if es.IsRegister {
		u, err := r.d.GetUserByEmail(req.Context(), sql.NullString{
			String: es.UserEmail,
		})
		if err != nil {
			logger.Error(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		_, err = r.d.UpdateUserIsDraft(req.Context(), data.UpdateUserIsDraftParams{IsDraft: false})
		if err != nil {
			logger.Error(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		logger.Debug(es.UserName, u.ID)
	}
	if u.IsDraft {

	}
	s, token, err := r.auth.NewAuthSession(int(u.ID), u.Email.String, es.UserName)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	_, err = r.auth.AddAuthSession(token, s)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = r.auth.UpdateIsExpiredSession(token)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	r.auth.SetCookie(w, "auth", token, time.Now().Add(time.Hour*192))

	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	r.auth.MagicLink.RemoveSession(dto.Email)
}

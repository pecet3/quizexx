package auth_router

import (
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
	var us data.User
	if es.IsRegister {
		u, err := r.d.GetUserByID(req.Context(), int64(es.UserID))
		if err != nil {
			logger.Error(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		us, err = r.d.UpdateUserIsDraft(req.Context(), data.UpdateUserIsDraftParams{IsDraft: false, ID: u.ID})
		if err != nil {
			logger.Error(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		logger.Debug(es.UserName, u.ID)
		logger.Debug(us)

	}
	logger.Debug(us)
	s, token, err := r.auth.NewAuthSession(us.ID, us.Email.String, es.UserName)
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
	r.auth.SetCookie(w, "refresh", s.RefreshToken, time.Now().Add(time.Hour*192))
	logger.Debug("refresh")
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	r.auth.MagicLink.RemoveSession(dto.Email)
}

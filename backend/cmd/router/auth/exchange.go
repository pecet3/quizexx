package auth_router

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleExchange(w http.ResponseWriter, req *http.Request) {
	dto := &dtos.ExchangeDTO{}
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

	if es.IsRegister {
		// create a new user
		u, err := r.d.User.GetByEmail(r.d.Db, es.UserEmail)
		if err != nil {
			logger.Error(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		u.IsDraft = false
		if err = u.Update(r.d.Db); err != nil {
			logger.Error(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		logger.Debug(es.UserName, u)
	}
	uDb, err := r.d.User.GetByEmail(r.d.Db, es.UserEmail)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	s, token, err := r.auth.NewAuthSession(uDb.ID, uDb.Email, es.UserName)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = r.auth.AddAuthSession(token, s)
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

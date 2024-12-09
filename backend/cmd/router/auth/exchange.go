package auth_router

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/data/entities"
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
	s, exists := r.auth.MagicLink.GetSession(dto.Email)
	if !exists {
		logger.WarnC("email sessions doesn't exist")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if s.ActivateCode != dto.Code {
		logger.WarnC("provided a wrong code ", s.UserEmail)
		logger.Debug(s.ActivateCode, dto.Code)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if s.ExchangeCounter > 5 || s.IsBlocked {
		s.IsBlocked = true
		logger.Warn("exchange counter block", s.UserEmail)
		http.Error(w, "Your account is blocked due the security reasons.", http.StatusBadRequest)
		return
	}
	s.ExchangeCounter += 1

	if s.IsRegister {
		// create a new user
		u := &entities.User{
			Name:      s.UserName,
			Email:     s.UserEmail,
			CreatedAt: time.Now(),
		}
		id, err := u.Add(r.d.Db)
		if err != nil {
			logger.Error(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		logger.Debug(s.UserName, id)
	}
	uDb, err := r.d.User.GetByEmail(r.d.Db, s.UserEmail)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	session, token, err := r.auth.NewAuthSession(uDb.ID, uDb.Email, s.UserName)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = r.auth.AddAuthSession(token, session)
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
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	r.auth.MagicLink.RemoveSession(dto.Email)
}

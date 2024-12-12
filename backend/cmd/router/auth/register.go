package auth_router

import (
	"encoding/json"
	"net/http"

	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/data/entities"
	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleRegister(w http.ResponseWriter, req *http.Request) {
	logger.Debug()
	dto := &dtos.RegisterDTO{}
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
	existingUser, _ := r.d.User.GetByEmail(r.d.Db, dto.Email)

	if existingUser != nil {
		if !existingUser.IsDraft {
			logger.Error(err)
			http.Error(w, "User with provided email already exists", http.StatusBadRequest)
			return
		}
	} else {
		u := &entities.User{}
		u.Email = dto.Email
		u.Name = dto.Name
		u.IsDraft = true
		_, err = u.Add(r.d.Db)
		if err != nil {
			logger.Error(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
	}

	logger.Debug(dto)
	s, code := r.auth.MagicLink.NewSessionRegister(dto.Name, dto.Email)
	if err = r.auth.MagicLink.AddSession(s); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = r.auth.MagicLink.SendEmailRegister(dto.Email, code, dto.Name)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	logger.Debug(code)
}

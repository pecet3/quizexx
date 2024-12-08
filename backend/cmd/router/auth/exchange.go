package auth_router

import (
	"encoding/json"
	"net/http"

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
		logger.Warn("email sessions deosn't exist")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if s.ActivateCode != dto.Code {
		logger.Warn("provided a wrong code")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if s.IsRegister {
		u := &entities.User{
			Name:  s.UserName,
			Email: s.UserEmail,
		}
		id, err := u.Add(r.d.Db)
		if err != nil {
			logger.Error(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		logger.Debug(s.UserName, id)
	}
	jwtToken, err := r.auth.JWT.GenerateJWT(s.UserEmail, s.UserName)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(jwtToken)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

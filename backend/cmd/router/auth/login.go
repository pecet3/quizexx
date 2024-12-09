package auth_router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleLogin(w http.ResponseWriter, req *http.Request) {
	dto := &dtos.LoginDTO{}
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
	u, err := r.d.User.GetByEmail(r.d.Db, dto.Email)
	if u == nil || err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	s, code := r.auth.MagicLink.NewSessionLogin(u.ID, dto.Email)
	if err = r.auth.MagicLink.AddSession(s); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = r.auth.MagicLink.SendEmailLogin(u.Email, code, u.Name)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	logger.InfoC(fmt.Sprintf(`user with email: %s is login. Access Code: %s`, u.Email, s.ActivateCode))
}

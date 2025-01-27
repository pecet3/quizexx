package auth_router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleLogin(w http.ResponseWriter, req *http.Request) {
	dto := &dtos.Login{}
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
	u, err := r.d.GetUserByEmail(req.Context(), sql.NullString{
		String: dto.Email,
		Valid:  true,
	})
	if u.ID == 0 || err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	s, code := r.auth.MagicLink.NewSessionLogin(int(u.ID), dto.Email)
	if err = r.auth.MagicLink.AddSession(s); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	logger.Debug(u.Email.String)
	err = r.auth.MagicLink.SendEmailLogin(string(u.Email.String), code, u.Name)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	logger.InfoC(fmt.Sprintf(`user with email: %s is login. Access Code: %s`, u.Email, s.ActivateCode))
}

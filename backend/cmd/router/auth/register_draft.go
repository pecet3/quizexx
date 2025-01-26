package auth_router

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleRegisterDraft(w http.ResponseWriter, req *http.Request) {
	logger.Debug()
	dto := &dtos.Register{}
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

	existingUser, err := r.d.GetUserByEmail(req.Context(), sql.NullString{String: dto.Email})

	if existingUser.ID != 0 || err == nil {
		if !existingUser.IsDraft {
			logger.Error(err)
			http.Error(w, "User with provided email already exists", http.StatusBadRequest)
			return
		}
	} else {
		_, err = r.d.AddUser(req.Context(), data.AddUserParams{
			Uuid: uuid.NewString(),
			Name: dto.Name,
			Email: sql.NullString{
				String: dto.Email,
			},
			IsDraft: true,
		})
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

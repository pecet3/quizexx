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

func (r router) handleRegister(w http.ResponseWriter, req *http.Request) {
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
	logger.Debug(dto)
	existingUser, err := r.d.GetUserByEmail(req.Context(), sql.NullString{String: dto.Email})
	var u data.User
	if existingUser.ID != 0 || err == nil {
		if !existingUser.IsDraft {
			logger.Error(err)
			http.Error(w, "User with provided email already exists", http.StatusBadRequest)
			return
		}
	} else {
		u, err = r.d.AddUser(req.Context(), data.AddUserParams{
			Salt: uuid.NewString(),
			Uuid: uuid.NewString(),
			Name: dto.Name,
			Email: sql.NullString{
				String: dto.Email,
				Valid:  true,
			},
			IsDraft:  true,
			ImageUrl: "/api/img/avatar.png",
		})
		if err != nil {
			logger.Error(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		_, err = r.d.AddGameUser(req.Context(), data.AddGameUserParams{
			UserID:     u.ID,
			Level:      0,
			Exp:        0,
			GamesWins:  0,
			RoundWins:  0,
			Percentage: 0,
		})
		if err != nil {
			logger.Error(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
	}

	logger.Debug(dto)
	s, code := r.auth.MagicLink.NewSessionRegister(int(u.ID), dto.Name, dto.Email)
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

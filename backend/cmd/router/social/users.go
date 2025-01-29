package social_router

import (
	"encoding/json"
	"net/http"

	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleGetTop50Users(w http.ResponseWriter, req *http.Request) {
	var dto []*dtos.User

	users, err := r.d.GetUsersSortedByLevel(req.Context(), data.GetUsersSortedByLevelParams{
		Limit:  50,
		Offset: 0,
	})
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		logger.Error(err)
		return
	}
	for _, user := range users {
		dto = append(dto, user.ToDto(r.d))
		logger.Debug(dto)
	}
	logger.Debug(dto, users)
	err = json.NewEncoder(w).Encode(dto)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		logger.Error(err)
		return
	}
}

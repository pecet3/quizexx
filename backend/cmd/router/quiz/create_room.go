package quiz_router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleCreateRoom(w http.ResponseWriter, req *http.Request) {
	u, err := r.auth.GetContextUser(req)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	if isExists := r.quiz.CheckUserHasRoom(0); isExists {
		logger.Warn(fmt.Sprintf(`user with id: %s wanted to create a room, when them room exists`, "0"))
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	dto := &dtos.Settings{}
	if err := json.NewDecoder(req.Body).Decode(dto); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	room := r.quiz.CreateRoom(*dto, u.ID)
	game, err := room.CreateGame()
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	go room.Run(r.quiz)
	logger.Debug(game.Content)
}

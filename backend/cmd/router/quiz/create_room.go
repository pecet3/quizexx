package quiz_router

import (
	"net/http"

	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleCreateRoom(w http.ResponseWriter, req *http.Request) {
	if isExists := r.quiz.CheckUserHasRoom(0); isExists {
		logger.InfoC("it exists ")
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	settings := dtos.Settings{
		Name:       "room1",
		GenContent: "test",
		MaxRounds:  "5",
		Difficulty: "1",
	}
	room := r.quiz.CreateRoom(settings, 0)
	game, err := room.CreateGame()
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	go room.Run(r.quiz)
	logger.Debug(game.Content)
}

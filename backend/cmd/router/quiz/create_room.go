package quiz_router

import (
	"net/http"

	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleCreateRoom(w http.ResponseWriter, req *http.Request) {

	settings := dtos.Settings{
		Name:       "room1",
		GenContent: "test",
		MaxRounds:  "5",
		Difficulty: "1",
	}
	room := r.quiz.CreateRoom(settings, 0)
	_, err := room.CreateGame()
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	go room.Run(r.quiz)
	logger.Debug(room)
}

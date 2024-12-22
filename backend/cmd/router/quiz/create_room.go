package quiz_router

import (
	"net/http"

	"github.com/pecet3/quizex/pkg/logger"
	"github.com/pecet3/quizex/pkg/quiz"
)

func (r router) handleCreateRoom(w http.ResponseWriter, req *http.Request) {

	settings := quiz.Settings{}
	room := r.quiz.CreateRoom(settings, 0)
	_, err := room.CreateGame()
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	logger.Debug(room)
}

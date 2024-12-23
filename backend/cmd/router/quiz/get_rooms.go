package quiz_router

import (
	"encoding/json"
	"net/http"

	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleGetRooms(w http.ResponseWriter, req *http.Request) {
	rooms := r.quiz.GetRoomsList()

	var res dtos.Rooms
	res.Rooms = rooms
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

}

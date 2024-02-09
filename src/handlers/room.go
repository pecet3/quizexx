package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pecet3/quizex/ws"
)

func GetRoomsHandler(m *ws.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomNames := m.GetRoomNamesList()

		response := struct {
			Rooms []string `json:"rooms"`
		}{Rooms: roomNames}

		jsonData, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(jsonData)
	}
}

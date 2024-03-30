package app

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pecet3/quizex/ws"
)

type quizHandler struct {
	app     *app
	manager ws.IManager
}

func (app *app) routeQuiz() {
	routeHandler := &quizHandler{
		app:     app,
		manager: &ws.Manager{},
	}
	app.mux.HandleFunc("/ws", app.mux.ServeHTTP)
	app.mux.HandleFunc("/quiz", routeHandler.handleWs)
}

var (
	upgrader = &websocket.Upgrader{
		CheckOrigin:     checkOrigin,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func checkOrigin(r *http.Request) bool {
	return true
}

func (h *quizHandler) handleWs(w http.ResponseWriter, req *http.Request) {
	m := h.manager

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
		return
	}

	roomName := req.URL.Query().Get("room")
	if roomName == "" {
		return
	}
	name := req.URL.Query().Get("name")
	if name == "" || name == "serwer" || name == "klient" {
		return
	}
	difficulty := req.URL.Query().Get("difficulty")
	maxRounds := req.URL.Query().Get("maxRounds")
	category := req.URL.Query().Get("category")
	newRoom := req.URL.Query().Get("new")

	if len(category) >= 32 || len(category) < 5 {
		conn.Close()
		return
	}

	settings := ws.Settings{
		Name:         roomName,
		GameCategory: category,
		Difficulty:   difficulty,
		MaxRounds:    maxRounds,
	}

	currentRoom := m.GetRoom(roomName)

	if currentRoom != nil {
		if newRoom == "true" {
			conn.Close()
			return
		}

		for roomClient := range currentRoom.clients {
			if name == roomClient.name {
				conn.Close()
				break
			}
		}
	}

	if currentRoom == nil {
		if newRoom == "true" {
			currentRoom = m.CreateRoom(settings)
			go currentRoom.Run(m)
		} else {
			conn.Close()
			return
		}
	}

	isSpectator := false

	if currentRoom.game.IsGame {
		isSpectator = true
	}

	client := &client{
		conn:        conn,
		receive:     make(chan []byte),
		room:        currentRoom,
		name:        name,
		answer:      -1,
		points:      0,
		isReady:     false,
		isSpectator: isSpectator,
		isAnswered:  false,
	}

	log.Printf("New connection, userName: %v connected to room: %v", name, roomName)
	currentRoom.join <- client
	defer func() { currentRoom.leave <- client }()
	go client.write()
	client.read()
}
func (h *quizHandler) handleHello() (w http.ResponseWriter, r *http.Request) {

}

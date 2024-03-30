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

func (app *app) routeQuiz(m *ws.Manager) {

	routeHandler := &quizHandler{
		app:     app,
		manager: &ws.Manager{},
	}
	routeHandler.manager = m
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

	m.HandleWs(conn, settings, newRoom, name)
}

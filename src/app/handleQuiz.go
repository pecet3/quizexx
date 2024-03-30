package app

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pecet3/quizex/ws"
)

type quizHandler struct {
	manager ws.IManager
}

func (app *app) routeQuiz(m *ws.Manager, mux *http.ServeMux) {
	manager := m.NewManager()
	routeHandler := &quizHandler{
		manager: manager,
	}
	log.Println(2)
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		routeHandler.serveWs(manager, w, r)
	})
	mux.HandleFunc("/hello", routeHandler.hello)

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
func (h *quizHandler) hello(w http.ResponseWriter, req *http.Request) {
	// Prosty tekst do zwrócenia
	message := "Hello, world!"

	// Ustawienie odpowiedzi HTTP z prostym tekstem
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Println("Error writing response:", err)
	}
}

func (h *quizHandler) serveWs(manager *ws.Manager, w http.ResponseWriter, req *http.Request) {
	m := manager
	log.Println(1)
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

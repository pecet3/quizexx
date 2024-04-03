package app

import (
	"log"
	"net/http"

	"github.com/pecet3/quizex/external"
	"github.com/pecet3/quizex/ws"
)

type quizHandler struct {
	manager  ws.IManager
	external external.IExternal
}

func (app *app) routeQuiz(mux *http.ServeMux, m *ws.Manager) {
	routeHandler := &quizHandler{
		manager:  m,
		external: app.external,
	}

	mux.HandleFunc("/ws", routeHandler.serveWS)
	mux.HandleFunc("/hello", routeHandler.hello)
}
func (h *quizHandler) hello(w http.ResponseWriter, req *http.Request) {
	message := "Hello, world!"

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Println("Error writing response:", err)
	}
}
func (h *quizHandler) serveWS(w http.ResponseWriter, req *http.Request) {
	h.manager.ServeWs(*h.external.NewExternalService(), w, req)
}

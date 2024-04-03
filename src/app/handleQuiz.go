package app

import (
	"log"
	"net/http"

	"github.com/pecet3/quizex/ws"
)

type quizHandler struct {
	manager ws.IManager
}

func (app *app) routeQuiz(mux *http.ServeMux) {
	manager := app.wsManager.NewManager()
	routeHandler := &quizHandler{
		manager: manager,
	}

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		manager.ServeWs(*app.external.NewExternalService(), w, r)
	})
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

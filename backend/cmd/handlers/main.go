package handlers

import (
	"net/http"

	"github.com/pecet3/quizex/pkg/data"
	"github.com/pecet3/quizex/pkg/ws"
)

type handlers struct {
	d   *data.Data
	wsm *ws.Manager
}

const PREFIX = "/v1"
const BASE_URL = "localhost:3000"

func Run(
	srv *http.ServeMux,
	d *data.Data,
	wsm *ws.Manager,
) {

	h := handlers{
		d:   d,
		wsm: wsm,
	}
	srv.HandleFunc(PREFIX+"/ws", h.handleQuiz)
	srv.HandleFunc(PREFIX+"/hello", h.hello)
	srv.Handle("/", http.FileServer(http.Dir("view")))

}

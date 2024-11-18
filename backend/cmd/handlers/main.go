package handlers

import (
	"net/http"

	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/pkg/auth"
	"github.com/pecet3/quizex/pkg/ws"
)

type handlers struct {
	d    *data.Data
	auth *auth.Auth
	wsm  *ws.Manager
}

const PREFIX = "/v1"
const BASE_URL = "localhost:3000"

func Run(
	srv *http.ServeMux,
	d *data.Data,
	auth *auth.Auth,
	wsm *ws.Manager,
) {

	h := handlers{
		d:    d,
		wsm:  wsm,
		auth: auth,
	}
	srv.HandleFunc(PREFIX+"/ws", h.handleQuiz)
	srv.HandleFunc(PREFIX+"/hello", h.hello)
	srv.Handle("/", http.FileServer(http.Dir("view")))

	srv.HandleFunc(PREFIX+"/auth", h.auth.HandleGoogleLogin)
}

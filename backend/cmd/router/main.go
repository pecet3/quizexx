package router

import (
	"net/http"

	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/pkg/auth"
	"github.com/pecet3/quizex/pkg/ws"
)

type router struct {
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

	r := router{
		d:    d,
		wsm:  wsm,
		auth: auth,
	}
	srv.HandleFunc(PREFIX+"/ws", r.handleQuiz)
	srv.HandleFunc(PREFIX+"/hello", r.hello)
	srv.Handle("/", http.FileServer(http.Dir("view")))

	srv.HandleFunc(PREFIX+"/auth", r.auth.HandleGoogleLogin)
	srv.HandleFunc("/google-callback", r.auth.HandleGoogleCallback)

}

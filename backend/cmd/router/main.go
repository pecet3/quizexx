package router

import (
	"net/http"

	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/repos"
	"github.com/pecet3/quizex/pkg/auth"
	"github.com/pecet3/quizex/pkg/logger"
	"github.com/pecet3/quizex/pkg/ws"
)

type router struct {
	d    *data.Data
	auth *auth.Auth
	wsm  *ws.Manager
	log  logger.Logger
}

const PREFIX = "/v1"
const BASE_URL = "localhost:3000"

func Run(
	app repos.App,
) {

	r := router{
		d:    app.Data,
		wsm:  app.Wsm,
		auth: app.Auth,
		log:  app.Logger,
	}
	app.Srv.HandleFunc(PREFIX+"/ws", r.handleQuiz)
	app.Srv.HandleFunc(PREFIX+"/hello", r.hello)
	app.Srv.Handle("/", http.FileServer(http.Dir("view")))

	app.Srv.HandleFunc(PREFIX+"/auth", r.auth.HandleOAuthLogin)
	app.Srv.HandleFunc(PREFIX+"/google-callback", r.handleGoogleCallback)

}

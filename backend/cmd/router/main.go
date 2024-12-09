package router

import (
	"net/http"

	authRouter "github.com/pecet3/quizex/cmd/router/auth"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/repos"
	"github.com/pecet3/quizex/pkg/auth"
	"github.com/pecet3/quizex/pkg/ws"
)

type router struct {
	d    *data.Data
	auth *auth.Auth
	wsm  *ws.Manager
}

const PREFIX = "/v1"

func Run(
	app repos.App,
) {

	r := router{
		d:    app.Data,
		wsm:  app.Wsm,
		auth: app.Auth,
	}
	authRouter.Run(app)

	app.Srv.HandleFunc(PREFIX+"/ws", r.handleQuiz)
	app.Srv.Handle(PREFIX+"/hello", r.auth.Authorize(r.hello))
	app.Srv.Handle("/", http.FileServer(http.Dir("view")))

}

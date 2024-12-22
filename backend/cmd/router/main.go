package router

import (
	"net/http"

	authRouter "github.com/pecet3/quizex/cmd/router/auth"
	quizRouter "github.com/pecet3/quizex/cmd/router/quiz"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/repos"
	"github.com/pecet3/quizex/pkg/auth"
	"github.com/pecet3/quizex/pkg/quiz"
)

type router struct {
	d    *data.Data
	auth *auth.Auth
	quiz *quiz.Manager
}

const PREFIX = "/api"

func Run(
	app repos.App,
) {

	r := router{
		d:    app.Data,
		quiz: app.Quiz,
		auth: app.Auth,
	}
	authRouter.Run(app)
	quizRouter.Run(app)
	app.Srv.Handle(PREFIX+"/hello", r.auth.Authorize(r.hello))
	app.Srv.Handle("/", http.FileServer(http.Dir("view")))
	app.Srv.Handle(PREFIX+"/", http.FileServer(http.Dir("img")))

	app.Srv.Handle(PREFIX+"/img/{fname}", r.auth.Authorize(r.handleImages))

}

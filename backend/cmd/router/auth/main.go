package auth_router

import (
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

const PREFIX = "/v1/auth"

func Run(
	app repos.App,
) {

	r := router{
		d:    app.Data,
		wsm:  app.Wsm,
		auth: app.Auth,
	}
	app.Srv.HandleFunc(PREFIX+"/magic-link/register", r.handleMagicLinkRegister)
}

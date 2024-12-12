package auth_router

import (
	"github.com/go-playground/validator/v10"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/repos"
	"github.com/pecet3/quizex/pkg/auth"
	"github.com/pecet3/quizex/pkg/ws"
)

type router struct {
	d    *data.Data
	auth *auth.Auth
	wsm  *ws.Manager
	v    *validator.Validate
}

const PREFIX = "/api/auth"

func Run(
	app repos.App,
) {

	r := router{
		d:    app.Data,
		wsm:  app.Wsm,
		auth: app.Auth,
		v:    app.Validator,
	}
	app.Srv.HandleFunc(PREFIX+"/register", r.handleRegister)
	app.Srv.HandleFunc(PREFIX+"/login", r.handleLogin)
	app.Srv.HandleFunc(PREFIX+"/exchange", r.handleExchange)
	app.Srv.HandleFunc(PREFIX+"/ping", r.handlePing)
}

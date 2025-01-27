package auth_router

import (
	"github.com/go-playground/validator/v10"
	"github.com/pecet3/quizex/cmd/repos"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/pkg/auth"
)

type router struct {
	d    *data.Queries
	auth *auth.Auth
	v    *validator.Validate
}

const PREFIX = "/api/auth"

func Run(
	app repos.App,
) {

	r := router{
		d:    app.Data,
		auth: app.Auth,
		v:    app.Validator,
	}
	app.Srv.HandleFunc(PREFIX+"/register", r.handleRegister)
	app.Srv.HandleFunc(PREFIX+"/login", r.handleLogin)
	app.Srv.HandleFunc(PREFIX+"/exchange", r.handleExchange)

	app.Srv.Handle(PREFIX+"/ping", r.auth.Authorize(r.handlePing))
}

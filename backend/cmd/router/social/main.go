package social_router

import (
	"github.com/go-playground/validator/v10"
	"github.com/pecet3/quizex/cmd/repos"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/pkg/auth"
	"github.com/pecet3/quizex/pkg/fetchers"
	"github.com/pecet3/quizex/pkg/quiz"
	"github.com/pecet3/quizex/pkg/social"
)

type router struct {
	d    *data.Queries
	auth *auth.Auth
	v    *validator.Validate
	quiz *quiz.Manager
	f    fetchers.Fetchers
	s    *social.Social
}

const PREFIX = "/api/social"

func Run(
	app repos.App,
) {
	r := router{
		d:    app.Data,
		auth: app.Auth,
		v:    app.Validator,
		quiz: app.Quiz,
		f:    app.Fetchers,
		s:    app.Social,
	}
	app.Srv.HandleFunc("GET "+PREFIX+"/fun-facts/latest", r.handleFunFact)
	app.Srv.HandleFunc("GET "+PREFIX+"/users", r.handleGetTop50Users)

}

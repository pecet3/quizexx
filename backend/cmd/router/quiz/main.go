package quiz_router

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pecet3/quizex/cmd/repos"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/pkg/auth"
	"github.com/pecet3/quizex/pkg/fetchers"
	"github.com/pecet3/quizex/pkg/logger"
	"github.com/pecet3/quizex/pkg/quiz"
)

type router struct {
	d    *data.Queries
	auth *auth.Auth
	v    *validator.Validate
	quiz *quiz.Manager
	f    fetchers.Fetchers
}

const PREFIX = "/api/quiz"

func Run(
	app repos.App,
) {
	r := router{
		d:    app.Data,
		auth: app.Auth,
		v:    app.Validator,
		quiz: app.Quiz,
		f:    app.Fetchers,
	}

	app.Srv.Handle(PREFIX+"/rooms/{name}", r.auth.Authorize(r.handleQuiz))

	app.Srv.Handle("POST "+PREFIX+"/rooms", r.auth.Authorize(r.handleCreateRoom))
	app.Srv.Handle("GET "+PREFIX+"/rooms", r.auth.Authorize(r.handleGetRooms))
}

func (r router) handleQuiz(w http.ResponseWriter, req *http.Request) {
	u, err := r.auth.GetContextUser(req)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	r.quiz.ServeQuiz(w, req, &u)
}

package quiz_router

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/repos"
	"github.com/pecet3/quizex/pkg/auth"
	"github.com/pecet3/quizex/pkg/quiz"
)

type router struct {
	d    *data.Data
	auth *auth.Auth
	v    *validator.Validate
	quiz *quiz.Manager
}

const PREFIX = "/api/auth"

func Run(
	app repos.App,
) {

	r := router{
		d:    app.Data,
		auth: app.Auth,
		v:    app.Validator,
		quiz: app.Quiz,
	}
	app.Srv.HandleFunc(PREFIX+"/game", r.handleQuiz)

	app.Srv.HandleFunc(PREFIX+"/create-room", r.handleCreateRoom)

}
func (r router) handleQuiz(w http.ResponseWriter, req *http.Request) {
	r.quiz.ServeWs(w, req)
}

package social_router

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pecet3/quizex/cmd/repos"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/pkg/auth"
	"github.com/pecet3/quizex/pkg/fetchers"
	"github.com/pecet3/quizex/pkg/logger"
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
	lvl, prc := r.s.CalculateLevelByExp(30000)
	logger.Debug(lvl, prc)
	app.Srv.HandleFunc("GET "+PREFIX+"/fun-facts/latest", r.handleFunFact)
}

func (r router) handleFunFact(w http.ResponseWriter, req *http.Request) {
	ff, err := r.d.GetCurrentFunFact(req.Context())
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		logger.Error(err)
		return
	}
	dto := ff.ToDTO(r.d)
	if err := dto.Send(w); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		logger.Error(err)
		return
	}
}

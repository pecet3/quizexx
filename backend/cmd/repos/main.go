package repos

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/pkg/auth"
	"github.com/pecet3/quizex/pkg/fetchers"
	"github.com/pecet3/quizex/pkg/quiz"
	"github.com/pecet3/quizex/pkg/social"
)

type App struct {
	Srv       *http.ServeMux
	Data      *data.Queries
	Auth      *auth.Auth
	Validator *validator.Validate
	Quiz      *quiz.Manager
	Fetchers  fetchers.Fetchers
	Social    *social.Social
}

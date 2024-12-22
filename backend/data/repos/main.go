package repos

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/pkg/auth"
	"github.com/pecet3/quizex/pkg/quiz"
)

type App struct {
	Srv       *http.ServeMux
	Data      *data.Data
	Auth      *auth.Auth
	Validator *validator.Validate
	Quiz      *quiz.Manager
}

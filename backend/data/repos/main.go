package repos

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/pkg/auth"
	"github.com/pecet3/quizex/pkg/ws"
)

type App struct {
	Srv       *http.ServeMux
	Data      *data.Data
	Auth      *auth.Auth
	Validator *validator.Validate
	Wsm       *ws.Manager
}
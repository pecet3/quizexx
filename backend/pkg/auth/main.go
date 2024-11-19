package auth

import (
	"github.com/pecet3/quizex/data"
	"golang.org/x/oauth2"
)

type Auth struct {
	statesMap    statesMap
	sessionsMap  sessionsMap
	oauth2Config *oauth2.Config
}

func New(d *data.Data) *Auth {
	return &Auth{
		statesMap:    newStatesMap(),
		sessionsMap:  newSessionMap(d),
		oauth2Config: newOAuthConfig(),
	}
}

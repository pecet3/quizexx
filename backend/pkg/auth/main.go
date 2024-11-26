package auth

import (
	"github.com/pecet3/quizex/data"
	"golang.org/x/oauth2"
)

type Auth struct {
	statesMap    *statesMap
	sessionsMap  *sessionsMap
	codesMap     *codesMap
	oauth2Config *oauth2.Config
	d            *data.Data
}

func New(d *data.Data) *Auth {
	return &Auth{
		statesMap:    newStatesMap(),
		sessionsMap:  newSessionMap(),
		codesMap:     newCodesMap(),
		oauth2Config: newOAuthConfig(),
		d:            d,
	}
}

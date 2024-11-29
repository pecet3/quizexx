package auth

import (
	"github.com/pecet3/quizex/data"
	"golang.org/x/oauth2"
)

type Auth struct {
	tmpMap       tempMobileSessions
	sessionsMap  sessionsMap
	oauth2Config *oauth2.Config
	d            *data.Data
}

func New(d *data.Data) *Auth {
	return &Auth{
		tmpMap:       newTempMobileSessions(),
		sessionsMap:  newSessionMap(),
		oauth2Config: newOAuthConfig(),
		d:            d,
	}
}

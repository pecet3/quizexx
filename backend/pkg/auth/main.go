package auth

import "golang.org/x/oauth2"

type Auth struct {
	statesMap    statesMap
	oauth2Config *oauth2.Config
}

func New() *Auth {
	return &Auth{
		statesMap:    newStatesMap(),
		oauth2Config: newOAuthConfig(),
	}
}

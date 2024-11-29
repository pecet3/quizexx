package auth

import (
	"github.com/google/uuid"
	"github.com/pecet3/quizex/data/entities"
)

type Code struct {
	DbUser     *entities.User
	PubCode    string
	SecretCode string
	JwtToken   string
	State      string
}

func generateCode() string {
	return uuid.NewString()
}

func (a *Auth) NewCode(pubCode string) *Code {
	secretCode := generateCode()
	state := generateState()
	return &Code{
		SecretCode: secretCode,
		PubCode:    pubCode,
		State:      state,
		JwtToken:   "",
		DbUser:     nil,
	}
}

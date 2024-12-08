package auth

import "github.com/pecet3/quizex/pkg/auth/magic_link"

type Auth struct {
	MagicLink *magic_link.MagicLink
}

func New() *Auth {
	return &Auth{
		MagicLink: magic_link.New(),
	}
}

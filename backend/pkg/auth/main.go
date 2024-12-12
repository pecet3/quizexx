package auth

import (
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/pkg/auth/magic_link"
)

type Auth struct {
	MagicLink *magic_link.MagicLink
	JWT       jwtServices
	d         *data.Data
}

func New(d *data.Data) *Auth {
	a := &Auth{
		MagicLink: magic_link.New(),
		JWT:       jwtServices{},
		d:         d,
	}
	go a.MagicLink.CleanUpExpiredSessions()
	return a
}

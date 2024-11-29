package auth

import (
	"github.com/pecet3/quizex/data/entities"
	"github.com/pecet3/quizex/pkg/logger"
	"golang.org/x/oauth2"
)

func (a *Auth) NewSession(user *entities.User, token *oauth2.Token) (*entities.Session, error) {
	exp := getExp()
	session := &entities.Session{
		UserID:       user.ID,
		Exp:          exp,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}
	return session, nil
}

func (a *Auth) AddSession(session *entities.Session) error {
	a.sessionsMap.set(session.AccessToken, session)

	err := session.Add(a.d.Db)
	if err != nil {
		return err
	}
	logger.Debug()
	return nil
}

func (a *Auth) GetSession(token string) (*entities.Session, bool) {
	session, exists := a.sessionsMap.get(token)
	if !exists {
		s, err := a.d.Session.GetByToken(a.d.Db, token)
		if err != nil || s == nil {
			return nil, false
		}
		return s, true
	}
	return session, true
}

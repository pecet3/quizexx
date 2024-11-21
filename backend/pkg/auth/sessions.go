package auth

import (
	"github.com/pecet3/quizex/data/entities"
	"github.com/pecet3/quizex/pkg/logger"
)

func (a *Auth) NewSession(user *entities.User) (*entities.Session, error) {
	exp := getExp()
	token, err := generateJWT(user, exp)
	if err != nil {
		return nil, err
	}
	session := &entities.Session{
		UserID: user.ID,
		Exp:    exp,
		Token:  token,
	}
	return session, nil
}

func (a *Auth) AddSession(session *entities.Session) error {
	a.sessionsMap.set(session.Token, session)

	err := session.Add(a.d.Db)
	if err != nil {
		return err
	}
	logger.Debug(a.sessionsMap.sessions)
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

package auth

import (
	"log"

	"github.com/pecet3/quizex/data/entities"
)

func (a *Auth) AddSession(session *entities.Session) error {
	a.sessionsMap.set(session.Token, session)
	err := a.d.Session.Add(a.d.Db)
	if err != nil {
		return err
	}
	return nil
}

func (a *Auth) GetSession(token string) (*entities.Session, bool) {
	session, exists := a.sessionsMap.get(token)
	if !exists {
		s, err := a.d.Session.GetByToken(a.d.Db, token)
		if err != nil || s == nil {
			log.Println(err, s)
			return nil, false
		}
		return s, true
	}
	return session, true
}

package auth

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/pecet3/quizex/data/entities"
)

const (
	SUSPEND_POST_SECONDS = 10
)

type AuthSessions = map[string]*entities.Session

func (as *Auth) NewAuthSession(userId int, uEmail, uName string) (*entities.Session, string, error) {
	expiresAt := time.Now().Add(168 * 4 * time.Hour)
	jwtToken, err := as.JWT.GenerateJWT(uEmail, uName)
	if err != nil {
		return nil, "", err
	}
	ea := &entities.Session{
		Token:             jwtToken,
		Expiry:            expiresAt,
		UserId:            userId,
		Email:             uEmail,
		ActivateCode:      jwtToken,
		PostSuspendExpiry: time.Now().Add(SUSPEND_POST_SECONDS * time.Second),
	}

	return ea, jwtToken, nil
}

func (as *Auth) GetAuthSession(token string) (*entities.Session, error) {
	return as.d.Session.GetByToken(as.d.Db, token)
}

func (as *Auth) AddAuthSession(token string, session *entities.Session) error {
	return as.d.Session.Add(as.d.Db, session)

}
func (as *Auth) UpdateIsExpiredSession(token string) error {
	return nil
}

func (as *Auth) GetAuthSessionFromRequest(r *http.Request) (*entities.Session, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return nil, err
	}
	sessionToken := cookie.Value
	var s *entities.Session
	s, err = as.GetAuthSession(sessionToken)
	if err != nil {
		log.Println("<Auth> Session doesn't exist")
		return nil, err
	}
	return s, nil
}

func (as *Auth) GetContextSession(r *http.Request) (*entities.Session, error) {
	ctx := r.Context()
	session, ok := ctx.Value(sessionContextKey).(*entities.Session)
	if !ok {
		return nil, errors.New("session not found in context")
	}
	return session, nil
}

func (as *Auth) GetContextUser(r *http.Request) (*entities.User, error) {
	ctx := r.Context()
	session, ok := ctx.Value(sessionContextKey).(*entities.Session)
	u, err := as.d.User.GetById(as.d.Db, session.UserId)
	if !ok {
		return nil, errors.New("session not found in context")
	}
	if err != nil {
		return nil, errors.New("not found in db")
	}
	return u, nil
}

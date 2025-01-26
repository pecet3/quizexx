package auth

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/entities"
)

const (
	SUSPEND_POST_SECONDS = 10
)

type AuthSessions = map[string]*entities.Session

func (as *Auth) NewAuthSession(userId int, uEmail, uName string) (data.AddSessionParams, string, error) {
	expiresAt := time.Now().Add(168 * 4 * time.Hour)
	jwtToken, err := as.JWT.GenerateJWT(uEmail, uName)
	if err != nil {
		return data.AddSessionParams{}, "", err
	}
	ea := data.AddSessionParams{
		Token:             jwtToken,
		Expiry:            expiresAt,
		UserID:            int64(userId),
		Email:             uEmail,
		ActivateCode:      jwtToken,
		PostSuspendExpiry: sql.NullTime{Time: time.Now().Add(SUSPEND_POST_SECONDS * time.Second)},
	}

	return ea, jwtToken, nil
}

func (as *Auth) GetAuthSession(token string) (data.Session, error) {
	return as.d.GetSessionByToken(context.Background(), token)
}

func (as *Auth) AddAuthSession(token string, session data.AddSessionParams) (data.Session, error) {
	return as.d.AddSession(context.Background(), session)
}
func (as *Auth) UpdateIsExpiredSession(token string) error {
	return nil
}

func (as *Auth) GetAuthSessionFromRequest(r *http.Request) (data.Session, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return data.Session{}, err
	}
	sessionToken := cookie.Value

	s, err := as.GetAuthSession(sessionToken)
	if err != nil {
		log.Println("<Auth> Session doesn't exist")
		return data.Session{}, err
	}
	return s, nil
}

func (as *Auth) GetContextSession(r *http.Request) (data.Session, error) {
	ctx := r.Context()
	session, ok := ctx.Value(sessionContextKey).(data.Session)
	if !ok {
		return data.Session{}, errors.New("session not found in context")
	}
	return session, nil
}

func (as *Auth) GetContextUser(r *http.Request) (data.User, error) {
	ctx := r.Context()
	session, ok := ctx.Value(sessionContextKey).(data.Session)
	if !ok {
		return data.User{}, errors.New("session not found in context")
	}
	u, err := as.d.GetUserByID(ctx, int64(session.UserID))

	if err != nil {
		return u, errors.New("not found in db")
	}
	return u, nil
}

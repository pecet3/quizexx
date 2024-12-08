package magic_link

import (
	"errors"
	"time"

	"github.com/pecet3/quizex/pkg/logger"
)

type EmailSession struct {
	UserEmail       string
	ActivateCode    string
	Expiry          time.Time
	IsRegister      bool
	LastNewSession  time.Time
	UserID          int
	UserName        string
	AttemptCounter  int
	ExchangeCounter int
}

type emailSessions = map[string]*EmailSession

func (ml *MagicLink) NewSessionLogin(
	uID int,
	email string) (*EmailSession, string) {
	expiresAt := time.Now().Add(time.Minute * 5)

	code := generateCode()
	ea := &EmailSession{
		Expiry:       expiresAt,
		ActivateCode: code,
		IsRegister:   false,
		UserID:       uID,
		UserEmail:    email,
	}
	return ea, code
}

func (ml *MagicLink) NewSessionRegister(
	name,
	email string) (*EmailSession, string) {
	expiresAt := time.Now().Add(time.Minute * 5)
	code := generateCode()
	ea := &EmailSession{
		Expiry:       expiresAt,
		ActivateCode: code,
		IsRegister:   true,
		UserID:       -1,
		UserName:     name,
		UserEmail:    email,
	}
	return ea, code
}

func (ml *MagicLink) GetSession(email string) (*EmailSession, bool) {
	ml.sMu.Lock()
	defer ml.sMu.Unlock()
	session, exists := ml.emailSessions[email]
	return session, exists
}

func (ml *MagicLink) AddSession(session *EmailSession) error {
	ml.sMu.Lock()
	defer ml.sMu.Unlock()
	es, exists := ml.emailSessions[session.UserEmail]
	if !exists {
		ml.emailSessions[session.UserEmail] = session
	} else {
		if es.AttemptCounter < 5 {
			es.AttemptCounter += 1
			es.LastNewSession = time.Now()
		} else {
			return errors.New("too much attempts")
		}
	}
	logger.Debug(ml.emailSessions)
	return nil
}

func (ml *MagicLink) RemoveSession(email string) {
	ml.sMu.Lock()
	defer ml.sMu.Unlock()
	delete(ml.emailSessions, email)
}

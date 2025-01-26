package magic_link

import (
	"errors"
	"fmt"
	"time"

	"github.com/pecet3/quizex/pkg/logger"
)

const ExpirySec = 60 * 5

type EmailSession struct {
	UserEmail       string
	ActivateCode    string
	Expiry          time.Time
	IsRegister      bool
	IsBlocked       bool
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
	expiresAt := time.Now().Add(time.Second * ExpirySec)

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
	uID int,
	name,
	email string) (*EmailSession, string) {
	expiresAt := time.Now().Add(time.Minute * ExpirySec)
	code := generateCode()
	ea := &EmailSession{
		Expiry:       expiresAt,
		ActivateCode: code,
		IsRegister:   true,
		UserID:       uID,
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
		return nil
	}
	if es.IsBlocked {
		errmsg := fmt.Sprintf(`Account is locked, left:%d minutes`, int(es.Expiry.Sub(time.Now()).Minutes()))
		return errors.New(errmsg)
	}
	if es.AttemptCounter >= 5 {
		es.IsBlocked = true
		es.Expiry = time.Now().Add(time.Minute * 60)
		errmsg := fmt.Sprintln(`Too many login attempts, Your account has been blocked for an hour`)

		return errors.New(errmsg)
	}
	es.AttemptCounter = es.AttemptCounter + 1
	es.ActivateCode = session.ActivateCode
	logger.Debug(es.AttemptCounter)
	es.LastNewSession = time.Now()

	logger.Debug(ml.emailSessions)
	return nil
}

func (ml *MagicLink) RemoveSession(email string) {
	ml.sMu.Lock()
	defer ml.sMu.Unlock()
	delete(ml.emailSessions, email)
}

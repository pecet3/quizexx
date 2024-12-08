package magic_link

import (
	"time"

	"github.com/google/uuid"
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
	newcode := uuid.NewString()
	expiresAt := time.Now().Add(time.Minute * 5)

	activateCode := "1234"
	ea := &EmailSession{
		Expiry:       expiresAt,
		ActivateCode: activateCode,
		IsRegister:   false,
		UserID:       uID,
		UserEmail:    email,
	}
	return ea, newcode
}

func (ml *MagicLink) NewSessionRegister(
	name,
	email string) (*EmailSession, string) {
	expiresAt := time.Now().Add(time.Minute * 5)
	activateCode := "1234"
	ea := &EmailSession{
		Expiry:       expiresAt,
		ActivateCode: activateCode,
		IsRegister:   true,
		UserID:       -1,
		UserName:     name,
		UserEmail:    email,
	}
	return ea, activateCode
}

func (ml *MagicLink) GetSession(code string) (*EmailSession, bool) {
	ml.sMu.Lock()
	defer ml.sMu.Unlock()
	session, exists := ml.emailSessions[code]
	return session, exists
}

func (ml *MagicLink) AddSession(code string, session *EmailSession) error {
	ml.sMu.Lock()
	defer ml.sMu.Unlock()
	es, exists := ml.emailSessions[code]
	if !exists {
		ml.emailSessions[code] = session
	} else {
		es.AttemptCounter = es.AttemptCounter + 1
		es.LastNewSession = time.Now()
	}
	logger.Debug(ml.emailSessions)
	return nil
}

func (ml *MagicLink) RemoveSession(email string) {
	ml.sMu.Lock()
	defer ml.sMu.Unlock()
	delete(ml.emailSessions, email)
}

// func (ml *MagicLink) VerifyNewEmailSession(uId int) bool {
// 	es, exists := ml.GetEmailSession(uId)
// 	if !exists {
// 		return true
// 	}
// 	log.Println(es.AttemptCounter)
// 	if es.AttemptCounter < 5 {

// 		return true
// 	} else {
// 		log.Println(3)
// 		if time.Since(es.LastNewSession).Minutes() > 30 {
// 			log.Println("since")
// 			return true
// 		}
// 	}

// 	return false
// }

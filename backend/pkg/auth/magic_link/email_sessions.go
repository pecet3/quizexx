package magic_link

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/pecet3/quizex/data/entities"
)

type emailSessions = map[string]*entities.Session

func (ml *MagicLink) NewEmailSession(
	userId int,
	email string,
	expiresAt time.Time) (*entities.Session, string) {
	newToken := uuid.NewString()

	hash := sha256.New()
	hash.Write([]byte(newToken))
	activateCode := hex.EncodeToString(hash.Sum(nil))
	ea := &entities.Session{
		Token:        newToken,
		Expiry:       expiresAt,
		ActivateCode: activateCode,
		UserId:       userId,
		Email:        email,
	}
	return ea, newToken
}

func (ml *MagicLink) GetEmailSession(token string) (*entities.Session, bool) {
	ml.sMu.Lock()
	defer ml.sMu.Unlock()
	session, exists := ml.emailSessions[token]
	return session, exists
}

func (ml *MagicLink) AddEmailSession(token string, session *entities.Session) error {
	ml.sMu.Lock()
	defer ml.sMu.Unlock()
	su, exists := ml.getTmpSection(session.UserId)
	log.Println("exists ", exists, session.UserId)
	if !exists {
		su = ml.newTmpSection(session.UserId)
		err := ml.addTmpSection(su)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		su.NewSessionCounter = su.NewSessionCounter + 1
		su.LastNewSession = time.Now()
		ml.updateTmpSection(su)
	}
	log.Println(su.NewSessionCounter, "counter")
	ml.emailSessions[token] = session
	return nil
}
func (ml *MagicLink) VerifyNewEmailSession(uId int) bool {
	su, exists := ml.getTmpSection(uId)
	if !exists {
		return true
	}
	log.Println(su.NewSessionCounter)
	if su.NewSessionCounter < 5 {

		return true
	} else {
		log.Println(3)
		if time.Since(su.LastNewSession).Minutes() > 30 {
			log.Println("since")
			return true
		}
	}

	return false
}

func (ml *MagicLink) RemoveEmailSession(token string) {
	ml.sMu.Lock()
	defer ml.sMu.Unlock()
	delete(ml.emailSessions, token)
}
func (ml *MagicLink) RemoveEmailSessionWithSessionUser(token string, uId int) error {
	ml.sMu.Lock()
	defer ml.sMu.Unlock()
	delete(ml.emailSessions, token)
	_, exists := ml.getTmpSection(uId)
	if !exists {
		return errors.New("no user in session users")
	}
	delete(ml.tmpSessions, uId)
	return nil
}

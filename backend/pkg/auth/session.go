package auth

import (
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/entities"
)

type sessions = map[string]*entities.Session
type sessionsMap struct {
	mu       sync.Mutex
	sessions sessions
	d        *data.Data
}

func newSessionMap(d *data.Data) sessionsMap {
	return sessionsMap{
		sessions: make(map[string]*entities.Session),
		d:        d,
	}
}

func (sm *sessionsMap) NewAuthSession(userId int) (*entities.Session, string) {
	expiresAt := time.Now().Add(168 * 4 * time.Hour)
	newToken := uuid.NewString()

	ea := &entities.Session{
		Token:  newToken,
		Expiry: expiresAt,
		UserId: userId,
	}

	return ea, newToken
}

func (sm *sessionsMap) GetAuthSession(token string) (*entities.Session, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	session, exists := sm.sessions[token]
	if !exists {
		s, err := sm.d.Session.GetByToken(sm.d.Db, token)
		if err != nil || s == nil {
			log.Println(err, s)
			return nil, false
		}
		return s, true
	}
	return session, true
}

func (sm *sessionsMap) AddAuthSession(token string, session *entities.Session) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.sessions[token] = session
	sm.d.Session = *session
	err := sm.d.Session.Add(sm.d.Db)
	if err != nil {
		log.Println(err)
		return nil
	}
	return nil
}
func (sm *sessionsMap) RemoveAuthSession(token string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.sessions, token)
}

package auth

import (
	"net/http"
	"sync"

	"github.com/pecet3/quizex/data/entities"
)

type sessions = map[string]*entities.Session
type sessionsMap struct {
	mu       sync.Mutex
	sessions sessions
}

func newSessionMap() sessionsMap {
	return sessionsMap{
		sessions: make(map[string]*entities.Session),
	}
}

func (sm *sessionsMap) get(token string) (*entities.Session, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	session, exists := sm.sessions[token]

	return session, exists
}

func (sm *sessionsMap) set(token string, session *entities.Session) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.sessions[token] = session
}
func (sm *sessionsMap) delete(token string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.sessions, token)
}

func setTokenCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   24 * 60 * 60, // 24 godziny
	})
}

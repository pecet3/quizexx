package auth

import (
	"net/http"
	"sync"

	"github.com/google/uuid"
)

func generateState() string {
	uuid := uuid.NewString()
	return uuid
}

type statesMap struct {
	mu     sync.Mutex
	states map[string]bool
}

func newStatesMap() statesMap {
	return statesMap{
		states: make(map[string]bool),
	}
}

func (s *statesMap) set(key string, value bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.states[key] = value
}

func (s *statesMap) get(key string) (bool, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, exists := s.states[key]
	return value, exists
}

func (s *statesMap) delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.states, key)
}

func (s *statesMap) has(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.states[key]
	return exists
}

func validateState(state string) error {
	return nil
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
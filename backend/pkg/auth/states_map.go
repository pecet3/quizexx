package auth

import (
	"sync"

	"github.com/google/uuid"
	"github.com/pecet3/quizex/data/entities"
)

type state = string

type pubCode = string

type Code struct {
	dbUser     *entities.User
	pubCode    string
	secretCode string
	jwtToken   string
}

func generateCode() string {
	uuid := uuid.NewString()
	return uuid
}

func (a *Auth) GetSecretCode(pubCode pubCode) string {
	secretCode := generateCode()
	code := &Code{
		secretCode: secretCode,
		jwtToken:   "",
	}
	a.tmpMap.set(pubCode, code)
	return secretCode
}

type tempMobileSessions struct {
	mu  sync.Mutex
	tms map[state]*Code
}

func newTempMobileSessions() tempMobileSessions {
	return tempMobileSessions{
		tms: make(map[state]*Code),
	}
}
func generateState() string {
	uuid := uuid.NewString()
	return uuid
}

func (s *tempMobileSessions) set(key string, value *Code) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tms[key] = value
}

func (s *tempMobileSessions) get(key string) (*Code, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, exists := s.tms[key]
	return value, exists
}

func (s *tempMobileSessions) delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.tms, key)
}

func (s *tempMobileSessions) has(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.tms[key]
	return exists
}

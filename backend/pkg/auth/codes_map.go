package auth

import (
	"sync"

	"github.com/google/uuid"
)

type pubCode = string

type code struct {
	secretCode string
	jwtToken   string
}
type codesMap struct {
	mu    sync.Mutex
	codes map[pubCode]code
}

func newCodesMap() *codesMap {
	return &codesMap{
		codes: make(map[pubCode]code),
	}
}
func generateCode() string {
	uuid := uuid.NewString()
	return uuid
}

func (a *Auth) GetSecretCode(pubCode pubCode) string {
	secretCode := generateCode()
	code := code{
		secretCode: secretCode,
		jwtToken:   "",
	}
	a.codesMap.set(pubCode, code)
	return secretCode
}

func (s *codesMap) set(key string, value code) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.codes[key] = value
}

func (s *codesMap) get(key string) (code, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, exists := s.codes[key]
	return value, exists
}

func (s *codesMap) delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.codes, key)
}

func (s *codesMap) has(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.codes[key]
	return exists
}

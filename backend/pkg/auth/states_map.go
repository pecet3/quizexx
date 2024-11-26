package auth

import (
	"sync"

	"github.com/google/uuid"
)

type state = string

type statesMap struct {
	mu     sync.Mutex
	states map[state]pubCode
}

func newStatesMap() *statesMap {
	return &statesMap{
		states: make(map[state]pubCode),
	}
}
func generateState() string {
	uuid := uuid.NewString()
	return uuid
}

func (s *statesMap) set(key string, value pubCode) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.states[key] = value
}

func (s *statesMap) get(key string) (pubCode, bool) {
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

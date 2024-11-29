package auth

import (
	"sync"

	"github.com/google/uuid"
)

// *********** public ***********
type tempMobileMap struct {
	mu  sync.RWMutex
	tms map[string]*Code
}

func (a *Auth) SetTmsCode(c *Code) {
	a.tmsMap.set(c.State, c)
}

func (a *Auth) GetTmsCode(state string) (*Code, bool) {
	return a.tmsMap.get(state)
}

//  ************ priv ****************

func newTempMobileMap() tempMobileMap {
	return tempMobileMap{
		tms: make(map[string]*Code),
	}
}

func generateState() string {
	return uuid.NewString()
}

func (s *tempMobileMap) set(key string, value *Code) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tms[key] = value
}

func (s *tempMobileMap) get(key string) (*Code, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, exists := s.tms[key]
	return value, exists
}

func (s *tempMobileMap) delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.tms, key)
}

func (s *tempMobileMap) has(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.tms[key]
	return exists
}

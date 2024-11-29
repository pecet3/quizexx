package auth

import (
	"sync"
)

// *********** public ***********
type jwtSmap struct {
	mu   sync.RWMutex
	jwts map[string]*Code
}

func (a *Auth) SetJWTCode(c *Code) {
	a.jwtSmap.set(c.PubCode, c)
}

func (a *Auth) GetJWTCode(pubCode string) (*Code, bool) {
	return a.jwtSmap.get(pubCode)
}

//  ************ priv ****************

func newjwtSmap() jwtSmap {
	return jwtSmap{
		jwts: make(map[string]*Code),
	}
}

func (s *jwtSmap) set(key string, value *Code) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jwts[key] = value
}

func (s *jwtSmap) get(key string) (*Code, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, exists := s.jwts[key]
	return value, exists
}

func (s *jwtSmap) delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.jwts, key)
}

func (s *jwtSmap) has(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.jwts[key]
	return exists
}

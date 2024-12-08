package magic_link

import (
	"sync"
)

type MagicLink struct {
	emailSessions emailSessions
	sMu           sync.RWMutex
	tmpSessions   tmpSessions
	tMu           sync.RWMutex
}

func New() *MagicLink {
	return &MagicLink{
		emailSessions: make(emailSessions),
		tmpSessions:   make(tmpSessions),
	}
}

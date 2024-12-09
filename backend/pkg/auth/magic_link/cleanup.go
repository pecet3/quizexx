package magic_link

import (
	"fmt"
	"time"

	"github.com/pecet3/quizex/pkg/logger"
)

func (ss *MagicLink) CleanUpExpiredSessions() {
	for {
		time.Sleep(60 * time.Second)
		cleanedEmailSessions := 0
		ss.sMu.Lock()
		for token, session := range ss.emailSessions {
			if time.Now().After(session.Expiry) {
				delete(ss.emailSessions, token)
				cleanedEmailSessions++
			}
		}
		ss.sMu.Unlock()
		logger.Info(fmt.Sprintf(`Cleaned Expired Email Sessions: %d`, cleanedEmailSessions))
	}
}

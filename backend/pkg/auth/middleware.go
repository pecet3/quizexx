package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/pecet3/quizex/data/entities"
	"github.com/pecet3/quizex/pkg/logger"
)

type contextKey string

const sessionContextKey contextKey = "session"

func (as *Auth) Authorize(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		sessionToken := cookie.Value
		var s *entities.Session
		s, err = as.GetAuthSession(sessionToken)
		if err != nil {
			logger.Warn("<Auth> Session doesn't exist")
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		if s.Expiry.Before(time.Now()) {
			as.MagicLink.RemoveSession("email")
			http.Error(w, "Your sessions is expired, you need to provide a new password", http.StatusUnauthorized)
			return
		}
		if r.Method == http.MethodPost {
			if !s.PostSuspendExpiry.IsZero() && !s.PostSuspendExpiry.Before(time.Now()) {
				logger.Warn("<Auth> User trying to use method POST, but they is suspended")
				http.Error(w, "suspended post method", http.StatusBadRequest)
				return
			}
			s.PostSuspendExpiry = time.Now().Add(10 * time.Second)
			as.d.Session.UpdatePostSuspendExpiry(as.d.Db, s.Token, s.PostSuspendExpiry)
		}
		ctx := context.WithValue(r.Context(), sessionContextKey, s)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

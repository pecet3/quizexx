package auth

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/pecet3/quizex/data/entities"
	"github.com/pecet3/quizex/pkg/logger"
)

type contextKey string

const sessionContextKey contextKey = "session"

func (as *Auth) Authorize(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "No authorization header", http.StatusUnauthorized)
			return
		}

		// Check if it's a Bearer token
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		// Extract the token
		jwt := headerParts[1]
		if jwt == "" {
			http.Error(w, "Empty token", http.StatusUnauthorized)
			return
		}
		_, err := as.JWT.ValidateJWT(jwt)
		if err != nil {
			logger.Warn("<Auth> Session doesn't exist")
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		var s *entities.Session
		s, err = as.GetAuthSession(jwt)
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

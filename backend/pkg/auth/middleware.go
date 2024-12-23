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

		var jwt string
		if r.Header.Get("Authorization") != "" {
			headerParts := strings.Split(r.Header.Get("Authorization"), " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}
			jwt = headerParts[1]
		} else {
			cookie, err := r.Cookie("auth")
			if err != nil || cookie.Value == "" {
				http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}
			jwt = cookie.Value
		}

		if jwt == "" {
			http.Error(w, "Empty token", http.StatusUnauthorized)
			return
		}
		claims, err := as.JWT.ValidateJWT(jwt)
		if err != nil {
			logger.Error(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		logger.Debug(claims)
		var s *entities.Session
		s, err = as.GetAuthSession(jwt)
		if err != nil || s == nil {
			logger.Warn("<Auth> Session doesn't exist")
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		if s.ActivateCode != jwt {
			logger.WarnC("invalid jwt token. user id: ", s.UserId)
			as.MagicLink.RemoveSession(s.Email)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		if s.Expiry.Before(time.Now()) {
			as.MagicLink.RemoveSession(s.Email)
			logger.Debug(s.Expiry.Before(time.Now()))
			http.Error(w, "Your sessions is expired you need to login once again", http.StatusUnauthorized)
			if err := as.UpdateIsExpiredSession(s.Token); err != nil {
				logger.Error(err)
				http.Error(w, "", http.StatusUnauthorized)
				return
			}
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

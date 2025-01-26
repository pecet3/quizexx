package auth

import (
	"context"
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/pkg/logger"
)

type contextKey string

const sessionContextKey contextKey = "session"

func (as *Auth) Authorize(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var jwt string
		var refresh string

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
			rCookie, err := r.Cookie("refresh")
			if err != nil || cookie.Value == "" {
				http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}
			refresh = rCookie.Value
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
		var s data.Session
		s, err = as.GetAuthSession(jwt)
		if err != nil || s.Token == "" {
			logger.Warn("<Auth> Session doesn't exist")
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		logger.Debug(s)
		if s.ActivateCode != jwt {
			logger.WarnC("invalid jwt token. user id: ", s.UserID)
			as.MagicLink.RemoveSession(s.Email)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		if s.Expiry.Before(time.Now()) {
			if refresh != s.RefreshToken {
				as.MagicLink.RemoveSession(s.Email)

				http.Error(w, "Your sessions is expired you need to login once again", http.StatusUnauthorized)
				if err := as.UpdateIsExpiredSession(s.Token); err != nil {
					logger.Error(err)
					http.Error(w, "", http.StatusUnauthorized)
					return
				}
				return
			}
		}
		logger.Debug(s.UserID)

		if r.Method == http.MethodPost {
			if !s.PostSuspendExpiry.Time.IsZero() && !s.PostSuspendExpiry.Time.Before(time.Now()) {
				logger.Warn("<Auth> User trying to use method POST, but they is suspended")
				http.Error(w, "suspended post method", http.StatusBadRequest)
				return
			}

			err := as.d.UpdatePostSuspendExpiry(
				r.Context(),
				data.UpdatePostSuspendExpiryParams{
					PostSuspendExpiry: sql.NullTime{Time: time.Now().Add(10 * time.Second)}})

			if err != nil {
				http.Error(w, "", http.StatusUnauthorized)
				return
			}

		}
		ctx := context.WithValue(r.Context(), sessionContextKey, s)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

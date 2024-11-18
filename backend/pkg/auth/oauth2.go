package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func newOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:9090/google-callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func (a *Auth) HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	state := generateState()
	a.statesMap.set(state, true)
	url := a.oauth2Config.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

func (a *Auth) HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	receivedState := r.URL.Query().Get("state")

	if isValid := a.statesMap.has(receivedState); !isValid {
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}
	a.statesMap.delete(receivedState)

	token, err := a.oauth2Config.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed to exc  hange token", http.StatusInternalServerError)
		return
	}

	user, err := getUserInfo(token.AccessToken)
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}

	jwtToken, err := generateJWT(user)
	if err != nil {
		http.Error(w, "Failed to generate JWT", http.StatusInternalServerError)
		return
	}

	setTokenCookie(w, jwtToken)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

}

// Pobranie informacji o użytkowniku z Google API
func getUserInfo(accessToken string) (*GoogleUser, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user GoogleUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Generowanie JWT
func generateJWT(user *GoogleUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// Middleware do autoryzacji
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Pobierz token z cookie
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Zweryfikuj token
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// claims := token.Claims.(jwt.MapClaims)
		// ctx := addUserToContext(r.Context(), claims)
		// next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Pomocnicze funkcje

package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func generateJWT(user *GoogleUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (a *Auth) ProcessJWT(user *GoogleUser, w http.ResponseWriter) error {
	jwtToken, err := generateJWT(user)
	if err != nil {
		return errors.New("failed to generate JWT")
	}
	a.sessionsMap.set()
	err = json.NewEncoder(w).Encode(jwtToken)
	if err != nil {
		return err
	}

	return nil
}

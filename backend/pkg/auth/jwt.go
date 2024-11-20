package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pecet3/quizex/data/entities"
)

func getExp() time.Time {
	return time.Now().Add(time.Hour * 24)
}

func generateJWT(user *entities.User, exp time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   exp.Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (a *Auth) ProcessJWT(user *entities.User, w http.ResponseWriter) error {
	exp := getExp()
	token, err := generateJWT(user, exp)
	if err != nil {
		return errors.New("failed to generate JWT")
	}
	session := &entities.Session{
		UserID: user.ID,
		Exp:    exp,
		Token:  token,
	}
	a.AddSession(session)
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		return err
	}

	return nil
}

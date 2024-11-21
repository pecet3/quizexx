package auth

import (
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

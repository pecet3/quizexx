package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("your_secret_key_here")

type jwtServices struct {
}

func (jwtServices) GenerateJWT(email, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    email,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token ważny przez 24 godziny
	})

	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (jwtServices) ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

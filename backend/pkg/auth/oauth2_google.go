package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/entities"
	"golang.org/x/oauth2"
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

func (gu *GoogleUser) ToDbUser(d *data.Data) *entities.User {
	userDb := &entities.User{
		Email:    gu.Email,
		Name:     fmt.Sprintf(`%s %s`, gu.Name, gu.FamilyName),
		ImageUrl: gu.Picture,
	}
	return userDb
}

func (a *Auth) GetGoogleUser(token *oauth2.Token) (*GoogleUser, error) {

	user, err := getUserInfo(token.AccessToken)
	if err != nil {
		return nil, errors.New("failed to get user info")
	}
	return user, nil
}
func (a *Auth) GetOAuth2Token(state, code string) (*oauth2.Token, error) {

	if isValid := a.tmpMap.has(state); !isValid {
		return nil, errors.New("invalid state")
	}
	a.tmpMap.delete(state)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	token, err := a.oauth2Config.Exchange(ctx, code)
	if err != nil {
		return nil, errors.New("failed to exchange token")
	}
	return token, nil
}

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

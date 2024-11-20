package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/entities"
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

func (a *Auth) GetGoogleUser(w http.ResponseWriter, r *http.Request) (*GoogleUser, error) {
	code := r.URL.Query().Get("code")
	receivedState := r.URL.Query().Get("state")

	if isValid := a.statesMap.has(receivedState); !isValid {
		return nil, errors.New("invalid state")
	}
	a.statesMap.delete(receivedState)

	token, err := a.oauth2Config.Exchange(r.Context(), code)
	if err != nil {
		return nil, errors.New("failed to exchange token")
	}

	user, err := getUserInfo(token.AccessToken)
	if err != nil {
		return nil, errors.New("failed to get user info")
	}
	return user, nil
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

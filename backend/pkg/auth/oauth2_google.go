package auth

import (
	"encoding/json"
	"net/http"
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

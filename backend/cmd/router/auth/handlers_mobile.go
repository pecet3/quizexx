package auth_router

import (
	"encoding/json"
	"net/http"

	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleMobileExchangeCodes(w http.ResponseWriter, req *http.Request) {
	logger.Debug("EXCHANGE CODES")
	queryParams := req.URL.Query()
	pubCode := queryParams.Get("pubCode")

	logger.Debug(pubCode)
	code := r.auth.NewCode(pubCode)
	r.auth.SetJWTCode(code)
	r.auth.SetTmsCode(code)
	secret := code.SecretCode
	err := json.NewEncoder(w).Encode(secret)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

}

func (r router) handleMobileAuth(w http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	pubCode := queryParams.Get("pubCode")
	if len(pubCode) <= 0 {
		logger.WarnC("no pubCode")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	code, exists := r.auth.GetJWTCode(pubCode)
	if !exists {
		logger.WarnC("code doesn't exist")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	logger.Debug(pubCode)
	url := r.auth.GetStateURL(code.State)
	http.Redirect(w, req, url, http.StatusTemporaryRedirect)
}

func (r router) handleMobileGoogleCallback(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	state := query.Get("state")
	code := query.Get("code")

	if state == "" || code == "" {
		logger.Error("Missing 'state' or 'code' in query parameters")
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	token, err := r.auth.GetOAuth2Token(state, code)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	gUser, err := r.auth.GetGoogleUser(token)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	dbUser := gUser.ToDbUser(r.d)
	logger.Debug(dbUser)
	session, err := r.auth.NewSession(dbUser, token)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	err = r.auth.AddSession(session)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	tmsCode, exists := r.auth.GetTmsCode(state)
	if !exists {
		logger.WarnC("no Code")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	tmsCode.JwtToken = token.AccessToken
	w.Write([]byte("<h1>You can go back to the app</h1>"))
}

func (r router) handleProvideJWT(w http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	pubCode := queryParams.Get("pubCode")
	secretCode := queryParams.Get("secretCode")
	if len(pubCode) <= 0 || len(secretCode) <= 0 {
		logger.WarnC("no pubCode or secretCode")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	code, exists := r.auth.GetJWTCode(pubCode)
	if !exists {
		logger.WarnC("no Code")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if code.SecretCode != secretCode {
		logger.WarnC("secret codes dont match")
		http.Error(w, "", http.StatusBadRequest)
		return
	}

}

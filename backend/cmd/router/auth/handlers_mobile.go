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
	secretCode := r.auth.GetSecretCode(pubCode)

	err := json.NewEncoder(w).Encode(secretCode)
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
	logger.Debug(pubCode)
	url := r.auth.GetStateURL(pubCode)
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
	session, err := r.auth.NewSession(dbUser)
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

}

package auth_router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleAuth(w http.ResponseWriter, req *http.Request) {
	logger.Debug("AUTH")
	queryParams := req.URL.Query()
	pubToken := queryParams.Get("pubToken")

	logger.Debug(pubToken)
	url := r.auth.GetStateURL()
	http.Redirect(w, req, url, http.StatusTemporaryRedirect)
}

func (r router) handleGoogleCallback(w http.ResponseWriter, req *http.Request) {
	gUser, err := r.auth.GetGoogleUser(w, req)
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
	log.Println(dbUser, gUser)

	// err = json.NewEncoder(w).Encode(session.Token)
	// if err != nil {
	// 	logger.Error(err)
	// 	http.Error(w, "", http.StatusUnauthorized)
	// 	return
	// }

	redirectUri := fmt.Sprintf("myapp://code?token=%s", session.Token)

	// Przekieruj użytkownika z powrotem do aplikacji mobilnej
	http.Redirect(w, req, redirectUri, http.StatusFound)
}

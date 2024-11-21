package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pecet3/quizex/pkg/logger"
)

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

	err = json.NewEncoder(w).Encode(session.Token)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

}

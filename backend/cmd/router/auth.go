package router

import (
	"log"
	"net/http"

	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleGoogleCallback(w http.ResponseWriter, req *http.Request) {
	user, err := r.auth.GetGoogleUser(w, req)
	if err != nil {
		logger.Error("google callback", err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	log.Println(user)
}

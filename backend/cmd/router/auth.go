package router

import (
	"encoding/json"
	"log"
	"net/http"
)

func (r router) handleGoogleCallback(w http.ResponseWriter, req *http.Request) {
	gUser, err := r.auth.GetGoogleUser(w, req)
	if err != nil {
		r.log.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	dbUser := gUser.ToDbUser(r.d)
	session, err := r.auth.NewSession(dbUser)
	if err != nil {
		r.log.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	err = r.auth.AddSession(session)
	if err != nil {
		r.log.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	err = json.NewEncoder(w).Encode(session.Token)
	if err != nil {
		r.log.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	log.Println(dbUser, gUser)
}

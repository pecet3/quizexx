package auth_router

import (
	"encoding/json"
	"net/http"

	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handlePing(w http.ResponseWriter, req *http.Request) {
	u, err := r.auth.GetContextUser(req)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

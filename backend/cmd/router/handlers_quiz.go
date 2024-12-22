package router

import (
	"log"
	"net/http"

	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) hello(w http.ResponseWriter, req *http.Request) {
	u, _ := r.auth.GetContextUser(req)
	logger.Debug(u)
	message := "Hello, world!"

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Println("Error writing response:", err)
	}
}

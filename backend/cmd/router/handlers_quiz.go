package router

import (
	"log"
	"net/http"
)

func (r router) hello(w http.ResponseWriter, req *http.Request) {
	message := "Hello, world!"

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Println("Error writing response:", err)
	}
}
func (r router) handleQuiz(w http.ResponseWriter, req *http.Request) {
	r.wsm.ServeWs(w, req)
}

package handlers

import (
	"log"
	"net/http"
)

func (h handlers) hello(w http.ResponseWriter, req *http.Request) {
	message := "Hello, world!"

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Println("Error writing response:", err)
	}
}
func (h handlers) handleQuiz(w http.ResponseWriter, req *http.Request) {
	h.wsm.ServeWs(w, req)
}

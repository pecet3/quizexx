package main

import (
	"log"
	"net/http"

	"github.com/pecet3/quizex/external"
	"github.com/pecet3/quizex/handlers"
	"github.com/pecet3/quizex/ws"
)

func main() {
	manager := ws.NewManager()

	mux := http.NewServeMux()

	mux.Handle("/ws", manager)

	mux.Handle("/", http.FileServer(http.Dir("view")))

	mux.HandleFunc("/api/rooms", handlers.GetRoomsHandler(manager))
	external.FetchBodyFromGPT("rock music", "easy", 5)
	address := "localhost:8080"
	log.Println("Server is running: ", address)
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

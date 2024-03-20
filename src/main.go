package main

import (
	"log"
	"net/http"

	"github.com/pecet3/quizex/handlers"
	"github.com/pecet3/quizex/ws"
)

func main() {
	manager := ws.NewManager()

	mux := http.NewServeMux()
	log.Println("Starting service")
	mux.Handle("/ws", manager)

	mux.Handle("/", http.FileServer(http.Dir("view")))

	mux.HandleFunc("/api/rooms", handlers.GetRoomsHandler(manager))
	address := "127.0.0.1:8090"
	log.Println("Server is running: ", address)
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

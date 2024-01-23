package main

import (
	"log"
	"net/http"

	"github.com/pecet3/quizex/ws"
)

func main() {
	manager := ws.NewManager()
	index := http.FileServer(http.Dir("view"))
	http.Handle("/", index)
	http.Handle("/ws", manager)

	address := "0.0.0.0:8080"

	server := &http.Server{
		Addr: address,
	}

	log.Println("Server is running: ", address)
	log.Fatal(server.ListenAndServe())
}

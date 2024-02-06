package main

import (
	"log"
	"net/http"

	"github.com/pecet3/quizex/ws"
)

func main() {
	manager := ws.NewManager()
	http.HandleFunc("/create", serveFile("view/create"))
	http.HandleFunc("/room", serveFile("view/room"))
	http.HandleFunc("/", serveFile("view"))

	address := "localhost:8080"
	log.Println("Server is running: ", address)
	http.Handle("/ws", manager)

	server := &http.Server{
		Addr: address,
	}

	log.Fatal(server.ListenAndServe())
}

func serveFile(directory string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, directory+r.URL.Path)
	}
}

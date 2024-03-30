package main

import (
	"log"

	"github.com/pecet3/quizex/app"
)

func main() {
	server := app.Run()

	log.Fatal(server.ListenAndServe())
}

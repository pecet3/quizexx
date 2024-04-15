package main

import (
	"log"

	"github.com/pecet3/quizex/app"
)

// I know I wrote shitty code but it works...
// Maybe I'd refactore it in the future but now I wanna be focused on the next thinks
func main() {
	log.Println("Running the server...")
	app.Run()
}

package main

import (
	"log"

	"github.com/pecet3/quizex/app"
)

// I know I wrote shitty code but it works... and I am aware it sucks
// I've tried understand how interfaces work and I know now that is no need to implement them here.
// Maybe I'd refactore it in the future but now I wanna be focused on the next projects

func main() {
	log.Println("Running the server...")
	app.Run()
}

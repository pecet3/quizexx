package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pecet3/quizex/pkg/logger"
)

const TEST = 10

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	go runAPI()

	<-stop
	onClose()
}
func onClose() {
	logger.Warn("Closing the server...")
}

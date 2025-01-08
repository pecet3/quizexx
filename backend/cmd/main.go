package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pecet3/quizex/pkg/logger"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	email := logger.Email{
		SMTPHost:     "smtp.example.com",
		SMTPPort:     587,
		Username:     "your-email@example.com",
		Password:     "your-password",
		FromAddress:  "your-email@example.com",
		ToAddresses:  []string{"recipient@example.com"},
		Subject:      "Test Email",
		BodyTemplate: "This is a test email sent using Go.",
	}

	loggerConfig := &logger.Config{IsDebugMode: true, Email: &email, Duration: time.Second * 10}
	logger := logger.New(loggerConfig)
	logger.Error("test")
	go runAPI()

	<-stop
	onClose()
}
func onClose() {
	logger.Warn("Closing the server...")
}

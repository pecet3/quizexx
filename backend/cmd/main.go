package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/pkg/logger"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go runAPI()
	test()
	<-stop
	onClose()
}
func onClose() {
	logger.Warn("Closing the server...")
}

func test() {
	sqle := data.New(data.NewSQLiteTest())

	u, err := sqle.GetUserByID(context.Background(), 1)
	logger.Debug(u, err)

	u, err = sqle.UpdateUserName(context.Background(), data.UpdateUserNameParams{
		Name: "kuba",
		ID:   u.ID,
	})
	logger.Debug(u, err)
}

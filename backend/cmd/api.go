package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator"
	"github.com/pecet3/quizex/cmd/handlers"
	"github.com/pecet3/quizex/pkg/auth"
	"github.com/pecet3/quizex/pkg/data"
	"github.com/pecet3/quizex/pkg/utils"
	"github.com/pecet3/quizex/pkg/ws"
)

type App struct {
	data *data.Data
	auth *auth.Auth
	v    *validator.Validate
	wsm  *ws.Manager
}

const BASE_URL = "localhost:5173"

func runAPI() {
	log.Println("Running the server...")
	utils.LoadEnv()

	mux := http.NewServeMux()

	app := App{
		data: data.New(),
		auth: auth.New(),
		v:    validator.New(),
		wsm:  ws.NewManager(),
	}

	handlers.Run(mux, app.data, app.auth, app.wsm)

	address := "localhost:9090"
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Server is listening on: [%s]", address)
		log.Fatal(server.ListenAndServe())
	}()
	<-stop
	onSrvClose()
}
func onSrvClose() {
	log.Println("Closing a server, removing cache files...")
	os.RemoveAll("s/")
}

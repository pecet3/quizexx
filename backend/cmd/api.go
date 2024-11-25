package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/pecet3/quizex/cmd/router"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/repos"
	"github.com/pecet3/quizex/pkg/auth"
	"github.com/pecet3/quizex/pkg/logger"
	"github.com/pecet3/quizex/pkg/utils"
	"github.com/pecet3/quizex/pkg/ws"
)

const BASE_URL = "localhost:9090"

func runAPI() {
	logger.Info("Starting...")
	utils.LoadEnv()

	mux := http.NewServeMux()
	data := data.New()

	app := repos.App{
		Srv:       mux,
		Data:      data,
		Auth:      auth.New(data),
		Validator: validator.New(),
		Wsm:       ws.NewManager(),
	}

	router.Run(app)

	address := "localhost:9090"
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	logger.Debug("to jest test lorem lorawedsad asd asd awd sass d")

	logger.Info(fmt.Sprintf("Server is listening on: [%s]", address))
	log.Fatal(server.ListenAndServe())

}

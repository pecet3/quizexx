package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pecet3/quizex/cmd/router"
	"github.com/pecet3/quizex/cmd/router/repos"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/pkg/auth"
	"github.com/pecet3/quizex/pkg/logger"
	"github.com/pecet3/quizex/pkg/quiz"
	"github.com/pecet3/quizex/pkg/utils"
)

const BASE_URL = "localhost:9090"

func runAPI() {
	logger.Info("Starting...")
	utils.LoadEnv()

	mux := http.NewServeMux()

	data := data.New(data.NewSQLite())

	app := repos.App{
		Srv:       mux,
		Data:      data,
		Auth:      auth.New(data),
		Validator: validator.New(),
		Quiz:      quiz.NewManager(data),
	}

	router.Run(app)

	server := &http.Server{
		Addr:    BASE_URL,
		Handler: mux,
	}
	logger.Info(fmt.Sprintf("Server is listening on: [%s]", BASE_URL))
	log.Fatal(server.ListenAndServe())

}

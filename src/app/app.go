package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/pecet3/quizex/database"
)

type app struct {
	db  *sql.DB
	mux *http.ServeMux
}

func Run() *http.Server {
	mux := http.NewServeMux()
	app := &app{
		db:  database.ConnectDb(),
		mux: http.NewServeMux(),
	}

	app.routeQuiz(mux)
	mux.Handle("/", http.FileServer(http.Dir("view")))

	address := "127.0.0.1:8090"
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}
	log.Println("Server is running: ", address)
	log.Fatal(server.ListenAndServe())
	return server
}

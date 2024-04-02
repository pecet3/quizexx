package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/pecet3/quizex/database"
	"github.com/pecet3/quizex/external"
	"github.com/pecet3/quizex/ws"
)

type app struct {
	db        *sql.DB
	mux       *http.ServeMux
	wsManager *ws.Manager
	external  *external.ExternalService
}

func Run() *http.Server {
	mux := http.NewServeMux()
	app := &app{
		db:        database.ConnectDb(),
		mux:       http.NewServeMux(),
		wsManager: &ws.Manager{},
		external:  &external.ExternalService{},
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

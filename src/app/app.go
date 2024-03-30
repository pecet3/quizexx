package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/pecet3/quizex/database"
	"github.com/pecet3/quizex/ws"
)

type app struct {
	db        *sql.DB
	mux       *http.ServeMux
	wsManager *ws.Manager
}

func Run() *http.Server {
	manager := ws.NewManager()

	mux := http.NewServeMux()
	app := &app{
		db:        database.ConnectDb(),
		mux:       http.NewServeMux(),
		wsManager: ws.NewManager(),
	}
	app.routeQuiz()

	log.Println("Starting service")
	mux.Handle("/ws", manager)
	mux.Handle("/", http.FileServer(http.Dir("view")))

	address := "127.0.0.1:8090"
	log.Println("Server is running: ", address)
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}
	return server
}

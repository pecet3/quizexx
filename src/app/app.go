package app

import (
	"log"
	"net/http"

	"github.com/pecet3/quizex/external"
	"github.com/pecet3/quizex/ws"
)

type app struct {
	// db        *sql.DB
	mux       *http.ServeMux
	wsManager ws.IManager
	external  external.IExternal
}

func Run() *http.Server {
	mux := http.NewServeMux()
	app := &app{
		// db:        database.ConnectDb(),
		mux:       mux,
		wsManager: &ws.Manager{},
		external:  &external.ExternalService{},
	}

	manager := app.wsManager.NewManager()

	app.routeQuiz(mux, manager)

	app.routeView(mux)

	address := "127.0.0.1:8090"
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}
	log.Printf("> > >   Server is running // %s   < < <", address)
	log.Fatal(server.ListenAndServe())
	return server
}

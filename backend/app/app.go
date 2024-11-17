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
	wsManager *ws.Manager
	external  external.IExternal
}

func Run() *http.Server {
	mux := http.NewServeMux()
	app := &app{
		// db:        database.ConnectDb(),
		mux:       mux,
		wsManager: ws.NewManager(),
		external:  &external.ExternalService{},
	}

	app.routeQuiz(mux, app.wsManager)

	app.routeView(mux)

	addrDev := "127.0.0.1:8090"
	addr := addrDev
	// addr := "0.0.0.0:8090"

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	log.Printf(">=>->[ Server is listening on: %s ]<-<=<", addr)
	log.Fatal(server.ListenAndServe())
	return server
}

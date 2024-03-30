package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/pecet3/quizex/database"
	"github.com/pecet3/quizex/ws"
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
	manager := ws.Manager{}
	m := manager.NewManager()

	log.Println(app)
	app.routeQuiz(m, mux)

	log.Println("Starting service")

	mux.Handle("/", http.FileServer(http.Dir("view")))

	// tmpl := template.Must(template.ParseGlob("view/*.html"))
	// mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	err := tmpl.Execute(w, 0)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// })

	address := "127.0.0.1:8090"
	log.Println("Server is running: ", address)
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}
	return server
}

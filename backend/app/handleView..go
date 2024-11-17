package app

import (
	"net/http"
)

type routerHandler struct {
}

func (app *app) routeView(mux *http.ServeMux) {

	mux.Handle("/", http.FileServer(http.Dir("view")))

}

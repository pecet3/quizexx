package router

import (
	"net/http"
)

func (r router) handleImages(w http.ResponseWriter, req *http.Request) {
	fName := req.PathValue("fname")
	http.ServeFile(w, req, "./cmd/img/"+fName)
}

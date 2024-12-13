package router

import (
	"net/http"

	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleImages(w http.ResponseWriter, req *http.Request) {
	fName := req.PathValue("fname")
	logger.Debug(fName)
	http.ServeFile(w, req, "./cmd/img/"+fName)
}

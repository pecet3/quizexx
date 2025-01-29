package social_router

import (
	"net/http"

	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleFunFact(w http.ResponseWriter, req *http.Request) {
	ff, err := r.d.GetCurrentFunFact(req.Context())
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		logger.Error(err)
		return
	}
	dto := ff.ToDTO(r.d)
	if err := dto.Send(w); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		logger.Error(err)
		return
	}
}

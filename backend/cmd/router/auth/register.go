package auth_router

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleMagicLinkRegister(w http.ResponseWriter, req *http.Request) {
	dto := &dtos.RegisterDTO{}
	err := json.NewDecoder(req.Body).Decode(dto)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
	}
	v := validator.New()
	err = dto.Validate(v)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
	}
	logger.Debug(dto)
}

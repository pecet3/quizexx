package quiz_router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleCreateRoom(w http.ResponseWriter, req *http.Request) {
	u, err := r.auth.GetContextUser(req)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	if isExists := r.quiz.CheckUserHasRoom(int(u.ID)); isExists {
		logger.Warn(fmt.Sprintf(`user with id: %s wanted to create a room, when them room exists`, "0"))
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	dto := &dtos.Settings{}
	if err := json.NewDecoder(req.Body).Decode(dto); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	err = dto.Validate(r.v)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if room := r.quiz.GetRoom(dto.Name); room != nil {
		http.Error(w, "Room with provided name already exists!", http.StatusForbidden)
		return
	}

	room := r.quiz.CreateRoom(dto.Name, int(u.ID))
	game, err := room.CreateGame(dto)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	go room.Run(r.quiz)

	logger.Debug(game.Content)

	gc, err := r.d.AddGameContents(req.Context(), data.AddGameContentsParams{
		Uuid:        game.UUID,
		MaxRounds:   int64(dto.MaxRounds),
		Category:    dto.GenContent,
		GenContent:  dto.GenContent,
		ContentJson: game.ContentJSON,
	})

	for i, round := range game.Content {
		gcr, err := r.d.AddGameContentRound(req.Context(), data.AddGameContentRoundParams{
			QuestionContent:    round.Question,
			Round:              int64(i + 1),
			CorrectAnswerIndex: int64(round.CorrectAnswer),
			GameContentID:      int64(gc.ID),
		})

		for i := 0; i < 4; i++ {
			isCorrect := false
			if round.CorrectAnswer == i {
				isCorrect = true
			}

			_, err = r.d.AddGameRound(req.Context(), data.AddGameRoundParams{
				IsCorrect:          isCorrect,
				GameContentRoundID: gcr.ID,
				Content:            round.Answers[i],
			})
		}
		if err != nil {
			logger.Error(err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

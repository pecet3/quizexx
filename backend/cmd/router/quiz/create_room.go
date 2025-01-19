package quiz_router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/data/entities"
	"github.com/pecet3/quizex/pkg/logger"
)

func (r router) handleCreateRoom(w http.ResponseWriter, req *http.Request) {
	u, err := r.auth.GetContextUser(req)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	if isExists := r.quiz.CheckUserHasRoom(u.ID); isExists {
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

	room := r.quiz.CreateRoom(dto, u.ID)
	game, err := room.CreateGame()
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	go room.Run(r.quiz)

	logger.Debug(game.Content)

	go func() {
		gc := entities.GameContent{
			UUID:        game.UUID,
			MaxRounds:   game.MaxRounds,
			Category:    game.Category,
			GenContent:  game.Category,
			ContentJSON: game.ContentJSON,
		}
		gcID, err := gc.Add(r.d.Db)

		for i, round := range game.Content {
			gcr := entities.GameContentRound{
				QuestionContent:    round.Question,
				Round:              i + 1,
				CorrectAnswerIndex: round.CorrectAnswer,
				GameContentID:      gcID,
			}
			gcrID, err := gcr.Add(r.d.Db)

			for i := 0; i < 4; i++ {
				isCorrect := false
				if round.CorrectAnswer == i {
					isCorrect = true
				}
				gca := entities.GameContentAnswer{
					IsCorrect:          isCorrect,
					GameContentRoundID: gcrID,
					Content:            round.Answers[i],
				}
				_, err = gca.Add(r.d.Db)
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
	}()

}

package social

import (
	"context"
	"math"

	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/pkg/logger"
)

func (s *Social) CalculateLevelByExp(exp float64) (level int, progress float64) {
	baseExp := 200.0 // Experience required for level 1
	currentExp := exp
	level = 1

	// Calculate level
	for {
		var expNeeded float64

		// Experience growth rate:
		// - 1/16 (6.25%) up to level 159
		// - 1/8 (12.5%) starting from level 160
		if level < 160 {
			expNeeded = baseExp * math.Pow(1.0625, float64(level-1)) // 1/16 growth (6.25%)
		} else {
			expNeeded = baseExp * math.Pow(1.0625, 159) * math.Pow(1.125, float64(level-160)) // 1/8 growth (12.5%) from level 160
		}

		// Check if experience is enough for the next level
		if currentExp >= expNeeded {
			currentExp -= expNeeded
			level++
		} else {
			// Calculate the progress progress for the current level
			progress = (currentExp / expNeeded) * 100
			break
		}
	}

	return level, progress
}
func (s *Social) CalculateExp(total int, difficulty string, isWinner bool) (exp float64) {
	// Base experience for easy difficulty
	baseExp := 1.

	// Difficulty multipliers
	var multiplier float64
	switch difficulty {
	case "easy":
		multiplier = 1.0
	case "medium":
		multiplier = 1.25
	case "hard":
		multiplier = 1.5625
	case "veryhard":
		multiplier = 1.953125
	default:
		multiplier = 1.0 // Default to easy if difficulty is invalid
	}

	// Calculate base experience based on total and difficulty
	exp = float64(total) * baseExp * multiplier

	// Double the experience if the player is the winner
	if isWinner {
		exp *= 2
	}

	return exp
}

func (s *Social) IncrementUserRecordsAfterFinishQuiz(uID int64, level int, isWin bool, exp, progress float64) error {
	ctx := context.Background()
	gm, err := s.d.GetGameUserByUserID(ctx, uID)
	if err != nil {
		return err
	}
	gw := gm.GamesWins
	if isWin {
		gw = +1
	}
	gu, _ := s.d.UpdateGameUser(ctx, data.UpdateGameUserParams{
		Level:     int64(level),
		Exp:       float64(exp),
		GamesWins: gw,
		RoundWins: 0,
		ID:        gm.ID,
		Progress:  progress,
	})
	logger.Debug(gu)
	return nil
}


func (r *Room) performRoundAction() error {

	isNextRound := r.game.checkIfShouldBeNextRound()
	indexCurrentContent := r.game.Content[r.game.State.Round-1]
	indexOkAnswr := indexCurrentContent.CorrectAnswer
	strOkAnswr := indexCurrentContent.Answers[indexOkAnswr]
	isEndGame := r.game.checkIfIsEndGame()
	if isEndGame {
		err := r.game.sendGameState()
		if err != nil {
			logger.Error("finish game err send game", err)
			return err
		}

		if err = r.sendServerMessage("The correct answer is: " + strOkAnswr); err != nil {
			logger.Error("finish game err", err)
			return err
		}
		time.Sleep(1800 * time.Millisecond)
		if err = r.sendServerMessage("It's finish the game"); err != nil {
			logger.Error("finish game err", err)
			return err
		}
		time.Sleep(1800 * time.Millisecond)
		winners := r.game.findWinner()
		winnersStr := strings.Join(winners, ", ")
		if len(winners) == 1 && len(winners) > 0 {
			if err = r.sendServerMessage("The game wins: " + winnersStr); err != nil {
				logger.Error("finish game err", err)
				return err
			}
		} else {
			if err = r.sendServerMessage("The game win: " + winnersStr); err != nil {
				logger.Error("finish game err", err)
				return err
			}
		}
		r.game.IsGame = false
		logger.Debug(winners)
		return err
	}

	if isNextRound {
		r.game.State.Round++
		var err error
		winnersStr := strings.Join(r.game.State.RoundWinners, ", ")
		if len(r.game.State.RoundWinners) == 0 {
			err = r.sendServerMessage("No one wins this round")
		}
		if len(r.game.State.RoundWinners) == 1 {
			err = r.sendServerMessage("This round wins " + winnersStr)
		}
		if len(r.game.State.RoundWinners) >= 2 {
			err = r.sendServerMessage("This round win: " + winnersStr)
		} else if len(r.game.State.RoundWinners) == len(r.game.Players) {
			for _, p := range r.game.Players {
				if p.points == 0 {
					return err
				}
				p.points -= 5
			}
			err = r.sendServerMessage("This round win everyone!")
		}
		if err != nil {
			logger.Error(err)
			return err
		}

		if !isEndGame {
			newState := r.game.newGameState(r.game.Content)
			r.game.State = newState
		}
		time.Sleep(1800 * time.Millisecond)
		err = r.sendServerMessage("The correct answer is: " + strOkAnswr)
		if err != nil {
			logger.Error(err)
			return err
		}

		time.Sleep(3000 * time.Millisecond)
		err = r.sendServerMessage("Round " + strconv.Itoa(r.game.State.Round) + " just started!")
		if err != nil {
			logger.Error(err)
			return err
		}

		for _, client := range r.game.Players {
			if client.isAnswered {
				client.isAnswered = false
			}
		}
		if err = r.game.sendGameState(); err != nil {
			logger.Error(err)
			return err
		}
	}
}

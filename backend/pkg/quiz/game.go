package quiz

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/logger"
)

type Game struct {
	UUID             string
	Room             *Room
	State            *GameState
	IsGame           bool
	Players          map[UUID]*Player
	Settings         *dtos.Settings
	ContentJSON      string
	Content          GameContent
	SecLeftForAnswer int
	SecLeftMu        sync.RWMutex
}

type Player struct {
	user       *data.User
	isReady    bool
	isAnswered bool
	answer     int
	points     int
	roundsWon  []uint
	lastActive time.Time
}

type GameState struct {
	Round           int            `json:"round"`
	Question        string         `json:"question"`
	Answers         []string       `json:"answers"`
	Actions         []*RoundAction `json:"actions"`
	Score           []*PlayerScore `json:"score"`
	PlayersAnswered []UUID         `json:"players_answered"`
	RoundWinners    []UUID         `json:"-"`
}

type RoundAction struct {
	UUID   string `json:"uuid"`
	Answer int    `json:"answer"`
	Round  int    `json:"round"`
}

type PlayerScore struct {
	User       *dtos.User `json:"user"`
	Points     int        `json:"points"`
	RoundsWon  []uint     `json:"rounds_won"`
	IsAnswered bool       `json:"is_answered"`
}

type GameContent = []RoundQuestion

type RoundQuestion struct {
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
	CorrectAnswer int      `json:"correct_answer"`
}

func (g *Game) updateSecLeftForAnswer(sec int) {
	g.SecLeftMu.Lock()
	defer g.SecLeftMu.Unlock()
	g.SecLeftForAnswer = sec

}
func (g *Game) getSecLeftForAnswer() int {
	g.SecLeftMu.Lock()
	defer g.SecLeftMu.Unlock()
	return g.SecLeftForAnswer
}

func (g *Game) getGameContent(s *dtos.Settings) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	options := "Options for this quiz:" +
		" category: " + s.GenContent +
		", diffuculty:" + s.Difficulty +
		", content language: " + s.Language
	prompt := "return json for quiz game with " +
		strconv.Itoa(s.MaxRounds) + " questions." +
		options +
		" You have to return correct struct. This is just array of objects. Nothing more, start struct: [{ question, 4x answers, correct_answer(index)}] "

	rawJSON := ""
	var err error
	if s.Name == "test" {
		rawJSON, _ = fakeFetchFromGPT()
	} else {
		rawJSON, err = fetchFromGPT(ctx, prompt)
	}
	if err != nil {
		return err
	}

	var content GameContent

	if err = json.Unmarshal([]byte(rawJSON), &content); err != nil {
		return err
	}
	g.Content = content
	g.ContentJSON = rawJSON

	return nil
}

func (g *Game) newScore(d *data.Queries) []*PlayerScore {
	var score []*PlayerScore
	for _, p := range g.Players {
		playerScore := &PlayerScore{
			User:      p.user.ToDto(d),
			Points:    p.points,
			RoundsWon: p.roundsWon,
		}
		score = append(score, playerScore)
	}

	return score
}

func (g *Game) newGameState(d *data.Queries, content []RoundQuestion) *GameState {
	logger.Debug("> New Game state in room: ", g.Room.Name)
	g.Content = content
	score := g.newScore(d)
	return &GameState{
		Round:           g.State.Round,
		Question:        g.Content[g.State.Round-1].Question,
		Answers:         g.Content[g.State.Round-1].Answers,
		Actions:         []*RoundAction{},
		Score:           score,
		PlayersAnswered: []UUID{},
		RoundWinners:    []UUID{},
	}
}
func (g *Game) checkIfAllPlayerAnswered() bool {
	playersInGame := len(g.Room.game.Players)
	playersFinished := len(g.State.PlayersAnswered)
	if playersFinished == playersInGame && playersInGame > 0 {
		return true
	}
	return false
}

func (g *Game) checkIfIsLastRound() bool {
	return g.State.Round == g.Settings.MaxRounds
}

func (g *Game) checkAnswer(p *Player, action *RoundAction) bool {
	isGoodAnswer := false
	if p.user.Uuid == action.UUID {
		p.answer = action.Answer
		if action.Answer == g.Content[g.State.Round-1].CorrectAnswer && !p.isAnswered {
			isGoodAnswer = true
		}
	}
	return isGoodAnswer
}

func (g *Game) toggleClientIsAnswered(p *Player, action *RoundAction) {
	if action.Answer >= 0 && !p.isAnswered {
		g.State.PlayersAnswered = append(g.State.PlayersAnswered, p.user.Uuid)
		p.isAnswered = true
	}
}
func (g *Game) findGameWinners() ([]string, []*Player) {
	highestScore := 0
	for _, c := range g.Players {
		if highestScore < c.points {
			highestScore = c.points
			continue
		}
	}
	winners := []string{}
	wp := []*Player{}
	if highestScore <= 0 {
		return winners, wp
	}
	for _, p := range g.Players {
		if highestScore == p.points {
			wp = append(wp, p)
			winners = append(winners, p.user.Name)
			continue
		}
	}
	return winners, wp
}

func (g *Game) countPoints() bool {
	isEveryoneAnsweredGood := false
	if len(g.State.RoundWinners) == len(g.Players) {
		for _, p := range g.Players {
			p.points += 5
		}
		isEveryoneAnsweredGood = true
		return isEveryoneAnsweredGood
	}
	for _, uuid := range g.State.RoundWinners {
		if p, ok := g.Players[uuid]; ok {
			p.points += 10
		}
	}
	return isEveryoneAnsweredGood
}

func (g *Game) getStrOkAnswer() string {
	indexCurrentContent := g.Content[g.State.Round-1]
	indexOkAnswr := indexCurrentContent.CorrectAnswer
	return indexCurrentContent.Answers[indexOkAnswr]
}

func (g *Game) performRound(m *Manager, hb *time.Ticker, isTimeout bool) error {
	isEveryoneAnswered := g.checkIfAllPlayerAnswered()
	isLastRound := g.checkIfIsLastRound()

	if isEveryoneAnswered || isTimeout {
		hb.Stop()
		strOkAnswr := g.getStrOkAnswer()
		var err error
		winnersStr := strings.Join(g.State.RoundWinners, ", ")
		logger.Debug(len(g.State.RoundWinners))
		logger.Debug(g.State.RoundWinners)
		if isEveryoneAnsweredGood := g.countPoints(); isEveryoneAnsweredGood {
			err = g.Room.sendServerMessage("This round win everyone!")
		} else if len(g.State.RoundWinners) == 0 {
			err = g.Room.sendServerMessage("No one wins this round")
		} else if len(g.State.RoundWinners) == 1 {
			err = g.Room.sendServerMessage("This round wins " + winnersStr)
		} else if len(g.State.RoundWinners) >= 2 && len(g.State.RoundWinners) == len(g.Players) {
			err = g.Room.sendServerMessage("This round win: " + winnersStr)
		}
		if err != nil {
			logger.Error(err)
			return err
		}

		g.updateSecLeftForAnswer(-1)
		if !isLastRound {
			g.State.Round++
		}
		newState := g.newGameState(m.d, g.Content)
		g.State = newState
		time.Sleep(TFR_SHORT_DURATION)
		err = g.Room.sendServerMessage("The correct answer is: " + strOkAnswr)
		if err != nil {
			logger.Error(err)
			return err
		}

		time.Sleep(TFR_LONG_DURATION)

		err = g.Room.sendServerMessage("Round " + strconv.Itoa(g.State.Round) + " just started!")
		if err != nil {
			logger.Error(err)
			return err
		}

		for _, client := range g.Players {
			client.isAnswered = false
		}
		if err = g.sendGameState(); err != nil {
			logger.Error(err)
			return err
		}

	}
	if isLastRound && isEveryoneAnswered || isLastRound && isTimeout {
		g.countPoints()
		g.sendGameState()
		if err := g.Room.sendServerMessage("It's finish the game"); err != nil {
			logger.Error("finish game err", err)
			return err
		}
		time.Sleep(TFR_LONG_DURATION)
		winners, wp := g.findGameWinners()
		ctx := context.Background()
		for _, p := range wp {
			m.d.AddGameWinner(ctx, data.AddGameWinnerParams{
				Points: int64(p.points),
				GameID: int64(0),
				UserID: p.user.ID,
			})
		}
		winnersStr := strings.Join(winners, ", ")
		if len(winners) == 1 && len(winners) > 0 {
			if err := g.Room.sendServerMessage("The game wins: " + winnersStr); err != nil {
				logger.Error("finish game err", err)
				return err
			}
		} else {
			if err := g.Room.sendServerMessage("The game win: " + winnersStr); err != nil {
				logger.Error("finish game err", err)
				return err
			}
		}
		g.IsGame = false
	} else {
		g.updateSecLeftForAnswer(g.Settings.SecForAnswer)
		hb.Reset(HEARTBEAT_DURATION)
	}

	return nil
}

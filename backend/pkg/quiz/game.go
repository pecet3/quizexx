package quiz

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/data/entities"
	"github.com/pecet3/quizex/pkg/logger"
)

type Game struct {
	UUID        string
	Room        *Room
	State       *GameState
	IsGame      bool
	Players     map[UUID]*Player
	Settings    *dtos.Settings
	ContentJSON string
	Content     GameContent
}

type Player struct {
	user        *entities.User
	isReady     bool
	isSpectator bool
	isAnswered  bool
	answer      int
	points      int
	roundsWon   []uint
	lastActive  time.Time
}

type GameState struct {
	Round           int           `json:"round"`
	Question        string        `json:"question"`
	Answers         []string      `json:"answers"`
	Actions         []RoundAction `json:"actions"`
	Score           []PlayerScore `json:"score"`
	PlayersAnswered []string      `json:"players_answered"`
	RoundWinners    []string      `json:"-"`
}

type RoundAction struct {
	UUID   string `json:"uuid"`
	Answer int    `json:"answer"`
	Round  int    `json:"round"`
}

type PlayerScore struct {
	User       *entities.User `json:"user"`
	Points     int            `json:"points"`
	RoundsWon  []uint         `json:"rounds_won"`
	IsAnswered bool           `json:"is_answered"`
}

type GameContent = []RoundQuestion

type RoundQuestion struct {
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
	CorrectAnswer int      `json:"correct_answer"`
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

func (r *Room) CreateGame(settings *dtos.Settings) (*Game, error) {
	logger.Info("Creating a game in room: ", r.Name)

	newGame := &Game{
		UUID:     uuid.NewString(),
		Room:     r,
		State:    &GameState{Round: 1},
		IsGame:   false,
		Players:  make(map[UUID]*Player),
		Settings: settings,
		Content:  nil,
	}
	if err := newGame.getGameContent(settings); err != nil {
		return nil, err
	}
	r.game = newGame
	return newGame, nil
}

func (g *Game) newScore() []PlayerScore {
	var score []PlayerScore

	for _, p := range g.Players {
		playerScore := PlayerScore{
			User:      p.user,
			Points:    p.points,
			RoundsWon: p.roundsWon,
		}
		score = append(score, playerScore)
	}

	return score
}

func (g *Game) newGameState(content []RoundQuestion) *GameState {
	logger.Debug("> New Game state in room: ", g.Room.Name)
	g.Content = content
	score := g.newScore()
	return &GameState{
		Round:           g.State.Round,
		Question:        g.Content[g.State.Round-1].Question,
		Answers:         g.Content[g.State.Round-1].Answers,
		Actions:         []RoundAction{},
		Score:           score,
		PlayersAnswered: []string{},
	}
}
func (g *Game) checkIfShouldBeNextRound() bool {
	playersInGame := len(g.Room.clients)
	playersFinished := len(g.State.PlayersAnswered)
	if playersFinished == playersInGame && playersInGame > 0 {
		return true
	}
	return false
}

func (g *Game) checkIfIsEndGame() bool {
	isEqualMaxAndCurrentRound := g.State.Round == g.Settings.MaxRounds
	isNextRound := g.checkIfShouldBeNextRound()

	if isEqualMaxAndCurrentRound && isNextRound {
		return true
	}
	return false
}

func (g *Game) checkAnswer(p *Player, action *RoundAction) bool {
	isGoodAnswer := false
	if p.user.UUID == action.UUID {
		p.answer = action.Answer
		if action.Answer == g.Content[g.State.Round-1].CorrectAnswer && !p.isAnswered {
			isGoodAnswer = true
		}
	}
	return isGoodAnswer
}

func (g *Game) toggleClientIsAnswered(p *Player, action *RoundAction) {
	if action.Answer >= 0 && !p.isAnswered {
		g.State.PlayersAnswered = append(g.State.PlayersAnswered, p.user.Name)
		p.isAnswered = true
	}
}
func (g *Game) findWinner() []string {
	highestScore := 0
	for _, c := range g.Room.game.Players {
		if highestScore < c.points {
			highestScore = c.points
			continue
		}
	}
	winners := []string{}
	for _, c := range g.Room.game.Players {
		if highestScore == c.points {
			winners = append(winners, c.user.Name)
			continue
		}
	}
	return winners
}

package quiz

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/pecet3/quizex/data/entities"
	"github.com/pecet3/quizex/pkg/logger"
)

type Game struct {
	Room       *Room
	State      *GameState
	IsGame     bool
	Players    map[UUID]*Client
	Category   string
	Difficulty string
	MaxRounds  int
	Content    []RoundQuestion
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
type RoundQuestion struct {
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
	CorrectAnswer int      `json:"correct_answer"`
}

func (r *Room) CreateGame() (*Game, error) {
	logger.Info("Creating a game in room: ", r.settings.Name)
	maxRoundsInt := r.settings.MaxRounds
	maxRoundsStr := strconv.Itoa(r.settings.MaxRounds)

	difficulty := r.settings.Difficulty
	category := r.settings.GenContent
	lang := r.settings.Language
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	content, err := fetchQuestionSet(ctx, category, maxRoundsStr, difficulty, lang)
	if err != nil {
		return nil, err
	}
	var questions []RoundQuestion

	err = json.Unmarshal([]byte(content), &questions)
	if err != nil {
		return nil, err
	}

	newGame := &Game{
		Room:       r,
		State:      &GameState{Round: 1},
		IsGame:     false,
		Players:    r.clients,
		Category:   r.settings.GenContent,
		Difficulty: r.settings.Difficulty,
		MaxRounds:  maxRoundsInt,
		Content:    questions,
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
	logger.Debug(playersFinished, playersInGame)
	if playersFinished == playersInGame && playersInGame > 0 {
		return true
	}
	return false
}

func (g *Game) checkIfIsEndGame() bool {
	isEqualMaxAndCurrentRound := g.State.Round == g.MaxRounds
	isNextRound := g.checkIfShouldBeNextRound()

	if isEqualMaxAndCurrentRound && isNextRound {
		return true
	}
	return false
}

func (g *Game) checkAnswer(c *Client, action *RoundAction) bool {
	isGoodAnswer := false
	if c.user.UUID == action.UUID {
		c.answer = action.Answer
		if action.Answer == g.Content[g.State.Round-1].CorrectAnswer && !c.isAnswered {
			isGoodAnswer = true
		}
	}
	return isGoodAnswer
}

func (g *Game) toggleClientIsAnswered(c *Client, action *RoundAction) {
	if action.Answer >= 0 && !c.isAnswered {
		g.State.PlayersAnswered = append(c.room.game.State.PlayersAnswered, c.name)
		c.isAnswered = true
	}
}

package quiz

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"
)

type Game struct {
	Room       *Room
	State      *GameState
	IsGame     bool
	Players    map[*Client]string
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
	Name   string `json:"name"`
	Answer int    `json:"answer"`
	Round  int    `json:"round"`
}

type PlayerScore struct {
	Name       string `json:"name"`
	Points     int    `json:"points"`
	RoundsWon  []uint `json:"rounds_won"`
	IsAnswered bool   `json:"is_answered"`
}
type RoundQuestion struct {
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
	CorrectAnswer int      `json:"correct_answer"`
}

func (r *Room) CreateGame() (*Game, error) {
	log.Println("> Creating a game in room: ", r.settings.Name)
	maxRoundStr := r.settings.MaxRounds
	maxRoundsInt, err := strconv.Atoi(maxRoundStr)
	if err != nil {
		return nil, err
	}
	difficulty := r.settings.Difficulty
	category := r.settings.GenContent
	lang := r.settings.Language
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	content, err := fetchQuestionSet(ctx, category, maxRoundStr, difficulty, lang)
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

func (g *Game) NewGameState(content []RoundQuestion) *GameState {
	log.Println("> New Game state in room: ", g.Room.name)
	g.Content = content
	score := g.NewScore()
	return &GameState{
		Round:           g.State.Round,
		Question:        g.Content[g.State.Round-1].Question,
		Answers:         g.Content[g.State.Round-1].Answers,
		Actions:         []RoundAction{},
		Score:           score,
		PlayersAnswered: []string{},
	}
}

func (g *Game) NewScore() []PlayerScore {
	var score []PlayerScore

	for p := range g.Players {
		playerScore := PlayerScore{
			Name:      p.name,
			Points:    p.points,
			RoundsWon: p.roundsWon,
		}
		score = append(score, playerScore)
	}

	return score
}

func (g *Game) CheckIfShouldBeNextRound() bool {
	playersInGame := len(g.Players)
	playersFinished := len(g.State.PlayersAnswered)
	if playersFinished == playersInGame && playersInGame > 0 {
		return true
	}
	return false
}

func (g *Game) CheckIfIsEndGame() bool {

	isEqualMaxAndCurrentRound := g.State.Round == g.MaxRounds
	isNextRound := g.CheckIfShouldBeNextRound()

	if isEqualMaxAndCurrentRound && isNextRound {
		return true
	}
	return false
}

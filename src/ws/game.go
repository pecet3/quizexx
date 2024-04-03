package ws

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/pecet3/quizex/external"
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
	PlayersFinished []string      `json:"playersFinished"`
}

type RoundAction struct {
	Name   string `json:"name"`
	Answer int    `json:"answer"`
	Round  int    `json:"round"`
}

type PlayerScore struct {
	Name      string `json:"name"`
	Points    int    `json:"points"`
	RoundsWon []uint `json:"roundsWon"`
}
type RoundQuestion struct {
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
	CorrectAnswer int      `json:"correctAnswer"`
}

func CreateGame(r *Room, external external.ExternalService) *Game {
	log.Println("> Creating a game in room: ", r.name)
	maxRoundStr := r.settings.MaxRounds
	maxRoundsInt, err := strconv.Atoi(r.settings.MaxRounds)
	if err != nil {
		return &Game{}
	}
	difficulty := r.settings.Difficulty
	category := r.settings.GameCategory
	lang := "polish"

	e := external.NewExternalService()

	content, err := e.FetchQuestionSet(category, maxRoundStr, difficulty, lang)

	if err != nil {
		return &Game{}
	}
	var questions []RoundQuestion

	err = json.Unmarshal([]byte(content), &questions)
	if err != nil {
		log.Println("error with unmarshal data")
		log.Println(err)
	}

	newGame := &Game{
		Room:       r,
		State:      &GameState{Round: 1},
		IsGame:     false,
		Players:    r.clients,
		Category:   r.settings.GameCategory,
		Difficulty: r.settings.Difficulty,
		MaxRounds:  maxRoundsInt,
		Content:    questions,
	}

	return newGame
}

func (g *Game) NewGameState(content []RoundQuestion) *GameState {
	g.Content = content
	score := g.NewScore()
	return &GameState{
		Round:           g.State.Round,
		Question:        g.Content[g.State.Round-1].Question,
		Answers:         g.Content[g.State.Round-1].Answers,
		Actions:         []RoundAction{},
		Score:           score,
		PlayersFinished: []string{},
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
	playersFinished := len(g.State.PlayersFinished)
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

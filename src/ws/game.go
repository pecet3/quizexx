package ws

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/pecet3/quizex/external"
)

type Game struct {
	Room       *room
	State      *GameState
	IsGame     bool
	Players    map[*client]string
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

type ResponseGPT struct {
	Category    string          `json:"category"`
	Difficulity string          `json:"difficulity"`
	Language    string          `json:"language"`
	Questions   []RoundQuestion `json:"questions"`
}

func (r *room) CreateGame() *Game {
	log.Println("creating a game")

	response, err := external.FetchBodyFromGPT(r.settings.gameCategory, r.settings.difficulty, r.settings.gameCategory)
	if err != nil {
		log.Println(err)
	}
	log.Println(response)
	var data ResponseGPT

	err = json.Unmarshal([]byte(response), &data)
	if err != nil {
		log.Println("error with unmarshal data")
	}
	maxRounds, err := strconv.Atoi(r.settings.maxRounds)

	if err != nil {
		maxRounds = 5
	}
	newGame := &Game{
		Room:       r,
		State:      &GameState{Round: 1},
		IsGame:     false,
		Players:    r.clients,
		Category:   r.settings.gameCategory,
		Difficulty: r.settings.difficulty,
		MaxRounds:  maxRounds,
		Content:    data.Questions,
	}

	r.game = newGame

	log.Println("new game: ", newGame.Content[0].Question)
	return newGame
}

func (g *Game) NewGameState() *GameState {
	if g.State.Round == g.MaxRounds {
		return &GameState{}
	}

	score := g.NewScore()
	log.Println("g.Content[g.State.Round-1].Question", g.Content[g.State.Round-1].Question)
	return &GameState{
		Round:    g.State.Round,
		Question: g.Content[g.State.Round-1].Question,
		Answers:  g.Content[g.State.Round-1].Answers,
		Actions:  []RoundAction{},
		Score:    score,
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

func (g *Game) CheckIfShouldBeNextRound() {
	playersInGame := len(g.Players)
	playersFinished := len(g.State.PlayersFinished)
	if playersFinished == playersInGame && playersInGame > 0 {
		g.State.Round++
		newState := g.NewGameState()
		g.State = newState
	}
}

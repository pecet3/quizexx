package ws

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/pecet3/quizex/external"
)

type Game struct {
	Room        *room
	State       *GameState
	IsGame      bool
	Players     map[*client]string
	mutex       sync.Mutex
	Category    string
	Difficulity string
	MaxRounds   string
	Content     []RoundQuestion
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

func (r *room) CreateGame(settings settingsGPT) *Game {
	log.Println("creating a game")

	response, err := external.FetchBodyFromGPT(settings.gameCategory, "easy", settings.gameCategory)
	if err != nil {
		log.Println(err)
	}

	data := ResponseGPT{}

	err = json.Unmarshal([]byte(response), &data)
	if err != nil {
		log.Println("error with unmarshal data")
	}

	newGame := &Game{
		Room:        r,
		State:       &GameState{Round: 1},
		IsGame:      false,
		Players:     r.clients,
		mutex:       sync.Mutex{},
		Category:    data.Category,
		Difficulity: data.Difficulity,
		MaxRounds:   settings.maxRounds,
		Content:     data.Questions,
	}

	r.game = newGame

	log.Println("new game: ", newGame.Content[0].Question)
	return newGame
}

func (g *Game) NewGameState() *GameState {
	g.mutex.Lock()
	defer g.mutex.Unlock()
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

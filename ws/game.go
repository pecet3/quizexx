package ws

import (
	"encoding/json"
	"log"
	"sync"
)

type Game struct {
	State    *GameState
	Content  []QandA
	Category string
	IsGame   bool
	Players  map[*client]bool
	mutex    sync.Mutex
}

type GameState struct {
	Round    int           `json:"round"`
	Question string        `json:"question"`
	Answers  []string      `json:"answers"`
	Actions  []RoundAction `json:"actions"`
	Score    []PlayerScore `json:"score"`
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

type QandA struct {
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
	CorrectAnswer int      `json:"correctAnswer"`
}

func (r *room) CreateGame() *Game {
	log.Println("creating a game")

	content := []QandA{
		{
			Question:      "test1",
			Answers:       []string{"a", "b", "c", "d"},
			CorrectAnswer: 2,
		},
		{
			Question:      "test2",
			Answers:       []string{"a", "b", "c", "d"},
			CorrectAnswer: 2,
		},
		{
			Question:      "test3",
			Answers:       []string{"a", "b", "c", "d"},
			CorrectAnswer: 2,
		},
		{
			Question:      "test4",
			Answers:       []string{"a", "b", "c", "d"},
			CorrectAnswer: 2,
		},
		{
			Question:      "test5",
			Answers:       []string{"a", "b", "c", "d"},
			CorrectAnswer: 2,
		},
	}

	newGame := &Game{
		State:    &GameState{Round: 1},
		Content:  content,
		Category: "",
		IsGame:   false,
		Players:  r.clients,
		mutex:    sync.Mutex{},
	}

	r.game = newGame

	log.Println("new game: ", newGame.Content[0].Question)
	return newGame
}

func (g *Game) NewGameState() *GameState {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	score := g.NewScore()

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

func (g *Game) GetGameState() *GameState {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	return g.State
}

func (g *Game) SendGameState(r *room) error {
	log.Println(g.Category, "category send game")
	g.mutex.Lock()
	defer g.mutex.Unlock()

	state := g.State
	stateBytes, err := json.Marshal(state)
	log.Println("Send Game State to the client: ", g.State)
	if err != nil {
		log.Println("Error marshaling game state:", err)
		return err
	}
	event := Event{
		Type:    "update_gamestate",
		Payload: stateBytes,
	}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		log.Println("Error marshaling game state:", err)
		return err
	}
	for client := range r.clients {
		if client == nil {
			return err
		}
		client.receive <- eventBytes
	}
	return nil
}

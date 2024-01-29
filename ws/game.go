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
	Round          int           `json:"round"`
	Question       string        `json:"question"`
	Answers        []string      `json:"answers"`
	Actions        []RoundAction `json:"actions"`
	ActionsHistory []RoundAction `json:"actionsHistory"`
}

type RoundAction struct {
	Name   string `json:"name"`
	Answer int    `json:"answer"`
	Points int    `json:"points"`
	Round  int    `json:"round"`
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
		State:    &GameState{},
		Content:  content,
		Category: "",
		IsGame:   false,
		Players:  make(map[*client]bool),
		mutex:    sync.Mutex{},
	}

	log.Println("new game: ", newGame)
	return newGame
}

func (g *Game) NewGameState() *GameState {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	if g.State == nil {
		return &GameState{}
	}
	return &GameState{
		Round:          g.State.Round + 1,
		Question:       g.Content[g.State.Round-1].Question,
		Answers:        g.Content[g.State.Round-1].Answers,
		Actions:        []RoundAction{},
		ActionsHistory: []RoundAction{},
	}
}

func (g *Game) GetGameState() *GameState {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	return g.State
}

func (g *Game) GetRoundActions() []RoundAction {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	var roundActions []RoundAction
	for client := range g.Players {
		action := RoundAction{
			Name:   client.name,
			Answer: client.answer,
			Points: client.points,
			Round:  client.round,
		}
		roundActions = append(g.State.Actions, action)
	}
	return roundActions
}

func (g *Game) GetActionsHistory() []RoundAction {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	var roundActions []RoundAction
	for client := range g.Players {
		action := RoundAction{
			Name:   client.name,
			Answer: client.answer,
			Points: client.points,
			Round:  client.round,
		}
		roundActions = append(g.State.ActionsHistory, action)
	}
	return roundActions
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

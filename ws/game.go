package ws

import (
	"encoding/json"
	"log"
	"sync"
)

type Game struct {
	State    *GameState
	Content  QandA
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

func (g *Game) CreateGame(category string) *Game {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if g.IsGame == true {
		return nil
	}
	state := g.CreateGameState()
	game := &Game{
		State:    state,
		Content:  QandA{},
		Category: category,
		IsGame:   false,
		Players:  make(map[*client]bool),
		mutex:    sync.Mutex{},
	}

	log.Println("new game: ", g.Category, game)
	return game
}

func (g *Game) CreateGameState() *GameState {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	return &GameState{
		Round:          0,
		Question:       "",
		Answers:        []string{""},
		Actions:        []RoundAction{},
		ActionsHistory: []RoundAction{},
	}
}

func (g *Game) GetGameState() *GameState {

	return g.State
}

func (g *Game) GetRoundActions() []RoundAction {

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

func (g *Game) SendGameState() error {
	state := g.GetGameState()
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
	for client := range g.Players {
		if client == nil {
			return err
		}
		client.receive <- eventBytes
	}
	return nil
}

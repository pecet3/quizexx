package ws

import (
	"encoding/json"
	"log"
)

type RoundAction struct {
	Name   string `json:"name"`
	Answer int    `json:"answer"`
	Points int    `json:"points"`
	Round  int    `json:"round"`
}

type GameState struct {
	IsGame         bool          `json:"isGame"`
	Category       string        `json:"category"`
	Round          int           `json:"round"`
	Question       string        `json:"question"`
	Answers        []string      `json:"answers"`
	Actions        []RoundAction `json:"actions"`
	ActionsHistory []RoundAction `json:"actionsHistory"`
}

type QandA struct {
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
	CorrectAnswer int      `json:"correctAnswer"`
}

func (r *room) CreateGame() *GameState {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.play == false {
		return nil
	}

	newGame := &GameState{}
	r.game = newGame
	return newGame
}

func (r *room) GetGameState() *GameState {
	actions := r.GetRoundActions()
	state := &GameState{
		IsGame:         r.play,
		Category:       r.game.Category,
		Round:          r.game.Round,
		Question:       r.body[r.round-1].Question,
		Answers:        r.body[r.round-1].Answers,
		Actions:        actions,
		ActionsHistory: actions,
	}
	return state
}

func (r *room) GetRoundActions() []RoundAction {
	var roundActions []RoundAction
	for client := range r.clients {
		action := RoundAction{
			Name:   client.name,
			Answer: client.answer,
			Points: client.points,
			Round:  client.round,
		}
		roundActions = append(roundActions, action)
	}
	return roundActions
}

func (r *room) GetActionsHistory() []RoundAction {
	var roundActions []RoundAction
	for client := range r.clients {
		action := RoundAction{
			Name:   client.name,
			Answer: client.answer,
			Points: client.points,
			Round:  client.round,
		}
		roundActions = append(roundActions, action)
	}
	return roundActions
}

func (r *room) SendGameState() error {
	state := r.GetGameState()
	stateBytes, err := json.Marshal(state)
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

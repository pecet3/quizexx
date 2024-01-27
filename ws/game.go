package ws

import (
	"encoding/json"
	"log"
)

type Player struct {
	Name   string `json:"name"`
	Answer int    `json:"answer"`
	Points int    `json:"points"`
	Round  int    `json:"round"`
}

type GameState struct {
	IsGame          bool      `json:"isGame"`
	Category        string    `json:"category"`
	Round           int       `json:"round"`
	Question        string    `json:"question"`
	Answers         []string  `json:"answers"`
	Players         []*Player `json:"players"`
	PrevRoundWinner []string  `json:"prevRoundWinner"`
}

type QandA struct {
	question      string
	answers       []string
	correctAnswer int
}

func (r *room) GetGameState() *GameState {
	players := r.GetPlayers()
	state := &GameState{
		IsGame:   true,
		Category: "test",
		Round:    r.round,
		Question: r.body[r.round-1].question,
		Answers:  r.body[r.round-1].answers,
		Players:  players,
	}
	return state
}

func (r *room) GetPlayers() []*Player {
	var players []*Player
	for client := range r.clients {
		player := &Player{
			Name:   client.name,
			Answer: client.answer,
			Points: client.points,
			Round:  client.round,
		}
		players = append(players, player)
	}
	return players
}

func (r *room) SendGameState() error {
	log.Println("sending game state")
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

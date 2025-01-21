package quiz

import (
	"encoding/json"
	"log"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type SendMessageEvent struct {
	UserName string `json:"user_name"`
	Message  string `json:"message"`
}

type WaitingPlayer struct {
	Name    string `json:"name"`
	IsReady bool   `json:"is_ready"`
}

type WaitingState struct {
	Players []WaitingPlayer `json:"players"`
}

type ServerMessage struct {
	Message string `json:"message"`
}

type PlayersAnswered struct {
	Players []string `json:"players"`
}

type ChatMessage struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Date    string `json:"date"`
}

func (g *Game) sendPlayersAnswered() error {
	eventBytes, err := marshalEventToBytes(g.State.PlayersAnswered, "players_answered")
	if err != nil {
		return err
	}
	for _, client := range g.Room.clients {
		if client == nil {
			return err
		}
		client.receive <- eventBytes
	}
	return nil
}

func (r *Room) sendTimeForAnswer(t int) error {
	eventBytes, err := marshalEventToBytes(t, "set_timer")
	if err != nil {
		return err
	}
	for _, client := range r.clients {
		if client == nil {
			return err
		}
		client.receive <- eventBytes
	}
	return nil
}
func (r *Room) sendIsFinish() error {
	eventBytes, err := marshalEventToBytes(true, "finish_game")
	if err != nil {
		return err
	}
	for _, client := range r.clients {
		if client == nil {
			return err
		}
		client.receive <- eventBytes
	}
	return nil
}

func (r *Room) sendSettings() error {
	eventBytes, err := marshalEventToBytes(r.game.Settings, "room_settings")
	if err != nil {
		return err
	}
	for _, client := range r.clients {
		if client == nil {
			return err
		}
		client.receive <- eventBytes
	}
	return nil
}

func (r *Room) sendReadyStatus() error {
	var readyClients []WaitingPlayer

	for _, c := range r.clients {
		RoomClient := WaitingPlayer{
			Name:    c.user.Name,
			IsReady: c.player.isReady,
		}
		readyClients = append(readyClients, RoomClient)
	}

	RoomMsg := WaitingState{
		Players: readyClients,
	}

	eventBytes, err := marshalEventToBytes(RoomMsg, "waiting_state")
	if err != nil {
		return err
	}
	for _, client := range r.clients {
		if client == nil {
			return err
		}
		client.receive <- eventBytes
	}
	return nil
}
func (r *Room) sendServerMessage(msg string) error {
	serverMsg := ServerMessage{
		Message: msg,
	}

	eventBytes, err := marshalEventToBytes(serverMsg, "server_message")
	if err != nil {
		return err
	}
	for _, client := range r.clients {
		if client == nil {
			return err
		}
		client.receive <- eventBytes
	}
	return nil
}

func (g *Game) sendGameState() error {
	eventBytes, err := marshalEventToBytes(*g.State, "update_gamestate")
	if err != nil {
		return err
	}
	for _, client := range g.Room.clients {
		if client == nil {
			return err
		}
		client.receive <- eventBytes
	}
	return nil
}

func marshalEventToBytes[T any](payload T, eventType string) ([]byte, error) {
	p := payload
	stateBytes, err := json.Marshal(p)
	if err != nil {
		log.Println("Error marshaling game state:", err)
		return nil, err
	}
	event := Event{
		Type:    eventType,
		Payload: stateBytes,
	}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		log.Println("Error marshaling game state:", err)
		return nil, err
	}
	return eventBytes, nil
}

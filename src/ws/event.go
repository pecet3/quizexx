package ws

import (
	"encoding/json"
	"log"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type SendMessageEvent struct {
	UserName string `json:"userName"`
	Message  string `json:"message"`
}

type ReadyClient struct {
	Name    string `json:"name"`
	IsReady bool   `json:"isReady"`
}

type ReadyStatus struct {
	Clients []ReadyClient `json:"clients"`
}

type ServerMessage struct {
	Message string `json:"message"`
}

func (r *Room) SendIsSpectator() error {
	eventBytes, err := MarshalEventToBytes[bool](true, "")
	if err != nil {
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
func (r *Room) SendIsFinish() error {
	eventBytes, err := MarshalEventToBytes[bool](true, "finish_game")
	if err != nil {
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

func (r *Room) SendSettings() error {
	eventBytes, err := MarshalEventToBytes[Settings](r.settings, "room_settings")
	if err != nil {
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

func (r *Room) SendReadyStatus() error {
	var readyClients []ReadyClient

	for c := range r.clients {
		roomClient := ReadyClient{
			Name:    c.name,
			IsReady: c.isReady,
		}
		readyClients = append(readyClients, roomClient)
	}

	roomMsg := ReadyStatus{
		Clients: readyClients,
	}

	eventBytes, err := MarshalEventToBytes[ReadyStatus](roomMsg, "ready_status")
	if err != nil {
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
func (r *Room) SendServerMessage(msg string) error {
	serverMsg := ServerMessage{
		Message: msg,
	}

	eventBytes, err := MarshalEventToBytes[ServerMessage](serverMsg, "server_message")
	if err != nil {
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

func (g *Game) SendGameState() error {
	eventBytes, err := MarshalEventToBytes[GameState](*g.State, "update_gamestate")
	if err != nil {
		return err
	}
	for client := range g.Room.clients {
		if client == nil {
			return err
		}
		client.receive <- eventBytes
	}
	return nil
}

func MarshalEventToBytes[T any](payload T, eventType string) ([]byte, error) {
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

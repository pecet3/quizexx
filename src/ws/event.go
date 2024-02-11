package ws

import (
	"encoding/json"
	"log"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *client) error

const (
	EventSendMessage = "send_message"
	EventSendAnswer  = "send_answer"
)

type SendMessageEvent struct {
	UserName string `json:"userName"`
	Message  string `json:"message"`
}

func (r *room) SendRoomMsg(msg string) error {
	var roomClients []RoomClient

	for c := range r.clients {
		roomClient := RoomClient{
			Name:    c.name,
			IsReady: c.isReady,
		}
		roomClients = append(roomClients, roomClient)
	}

	roomMsg := RoomMsg{
		Message: msg,
		Clients: roomClients,
	}

	eventBytes, err := MarshalEventToBytes[RoomMsg](roomMsg)
	if err != nil {
		return err
	}
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
	log.Println(g.Category, "category send game")
	g.mutex.Lock()
	defer g.mutex.Unlock()

	eventBytes, err := MarshalEventToBytes[GameState](*g.State)
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

func MarshalEventToBytes[T any](payload T) ([]byte, error) {
	p := payload
	stateBytes, err := json.Marshal(p)
	log.Println("Send Game State to the client: ", p)
	if err != nil {
		log.Println("Error marshaling game state:", err)
		return nil, err
	}
	event := Event{
		Type:    "update_gamestate",
		Payload: stateBytes,
	}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		log.Println("Error marshaling game state:", err)
		return nil, err
	}
	return eventBytes, nil
}

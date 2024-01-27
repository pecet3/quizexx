package ws

import (
	"encoding/json"
	"errors"
	"log"
)

type room struct {
	name    string
	events  map[string]EventHandler
	clients map[*client]bool

	join  chan *client
	ready chan *client
	leave chan *client

	forward       chan []byte
	receiveAnswer chan []byte

	body           []QandA
	round          int
	playersActions []Player
	play           bool
}

func NewRoom(name string) *room {
	body := []QandA{
		{
			question:      "test1",
			answers:       []string{"a", "b", "c", "d"},
			correctAnswer: 2,
		},
		{
			question:      "test",
			answers:       []string{"a", "a", "c", "d"},
			correctAnswer: 1,
		},
	}
	r := &room{
		name:          name,
		clients:       make(map[*client]bool),
		join:          make(chan *client),
		leave:         make(chan *client),
		ready:         make(chan *client),
		forward:       make(chan []byte),
		receiveAnswer: make(chan []byte),
		body:          body,
		round:         1,
		events:        make(map[string]EventHandler),
		play:          false,
	}
	r.setupEventHandlers()
	return r
}

func (m *manager) GetRoom(name string) *room {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.rooms[name]
}

func (m *manager) CreateRoom(name string) *room {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if existingRoom, ok := m.rooms[name]; ok {
		return existingRoom
	}

	newRoom := NewRoom(name)
	m.rooms[name] = newRoom
	log.Println("Created a room with name:", name)
	return newRoom
}

func (m *manager) RemoveRoom(name string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if room, ok := m.rooms[name]; ok {
		for client := range room.clients {
			room.leave <- client
		}
		close(room.join)
		close(room.forward)
		close(room.ready)
		close(room.receiveAnswer)
		close(room.leave)

		delete(m.rooms, name)
		log.Println("Closing a room with name:", room.name)
		return
	}
}

func (r *room) Run(m *manager) {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true

		case client := <-r.leave:
			if len(r.clients) == 0 {
				log.Println("leaving")
				close(client.receive)
				m.RemoveRoom(r.name)
				return
			}
		case msg := <-r.forward:
			for client := range r.clients {
				client.receive <- msg
			}
		case ready := <-r.ready:
			ready.isReady = true
			r.SendGameState()

		case action := <-r.receiveAnswer:
			var player Player
			err := json.Unmarshal(action, &player)
			if err != nil {
				log.Println("Error marshaling game state:", err)
				return
			}

			r.playersActions = append(r.playersActions, player)
			if player.Answer == r.body[r.round-1].correctAnswer {
				for client := range r.clients {
					if client == nil {
						continue
					}
					if client.name == player.Name {
						client.points = client.points + 10
					}
				}
			}
			if len(r.clients) == len(r.playersActions) {
				r.round = r.round + 1
			}
			r.SendGameState()

		}

	}
}

func (r *room) setupEventHandlers() {
	r.events[EventSendMessage] = SendMessage
}

func SendMessage(event Event, c *client) error {
	return nil
}

func (r *room) routeEvent(event Event, c *client) error {
	if handler, ok := r.events[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("There is no such event type")
	}
}

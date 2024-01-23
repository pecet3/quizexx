package ws

import (
	"errors"
	"log"
)

type QandA struct {
	questions     []string
	correctAnswer int
}
type room struct {
	name    string
	events  map[string]EventHandler
	clients map[*client]bool

	join  chan *client
	ready chan *client
	leave chan *client

	forward chan []byte

	body  []QandA
	round chan int
	play  chan bool
}

func NewRoom(name string) *room {
	body := []QandA{
		{
			questions:     []string{"a", "b", "c", "d"},
			correctAnswer: 2,
		},
		{
			questions:     []string{"a", "a", "c", "d"},
			correctAnswer: 1,
		},
	}
	r := &room{
		name:    name,
		clients: make(map[*client]bool),
		join:    make(chan *client),
		leave:   make(chan *client),
		ready:   make(chan *client),
		forward: make(chan []byte),
		body:    body,
		round:   make(chan int),
		events:  make(map[string]EventHandler),
		play:    make(chan bool),
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
		close(room.join)
		close(room.forward)
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

			close(client.receive)

			if len(r.clients) == 0 {
				m.RemoveRoom(r.name)
				return
			}
		case msg := <-r.forward:
			for client := range r.clients {
				client.receive <- msg
			}

		case ready := <-r.ready:
			if len(r.ready) == len(r.clients) {
				r.round <- 1
				ready.round <- 1
				r.play <- true
			}
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

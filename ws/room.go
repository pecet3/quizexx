package ws

import (
	"encoding/json"
	"log"
	"sync"
)

type room struct {
	name     string
	category string
	clients  map[*client]bool
	mutex    sync.Mutex
	join     chan *client
	ready    chan *client
	leave    chan *client

	forward       chan []byte
	receiveAnswer chan []byte

	game         *GameState
	body         []QandA
	round        int
	roundActions []RoundAction
	play         bool
}

func NewRoom(name string) *room {
	body := []QandA{
		{
			Question:      "test1",
			Answers:       []string{"a", "b", "c", "d"},
			CorrectAnswer: 2,
		},
		{
			Question:      "test",
			Answers:       []string{"a", "a", "c", "d"},
			CorrectAnswer: 2,
		},
		{
			Question:      "test1",
			Answers:       []string{"a", "b", "c", "d"},
			CorrectAnswer: 2,
		},
		{
			Question:      "test1",
			Answers:       []string{"a", "b", "c", "d"},
			CorrectAnswer: 2,
		},
		{
			Question:      "test1",
			Answers:       []string{"a", "b", "c", "d"},
			CorrectAnswer: 2,
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
		play:          false,
	}

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
		case msg := <-r.forward:
			for client := range r.clients {
				client.receive <- msg
			}
		case client := <-r.join:
			r.clients[client] = true

		case client := <-r.leave:
			close(client.receive)
			delete(r.clients, client)

			if len(r.clients) == 0 {
				log.Println("leaving")
				m.RemoveRoom(r.name)
				return
			}

		case client := <-r.ready:
			client.isReady = true
			r.play = true
			r.SendGameState()

		case action := <-r.receiveAnswer:
			var actionPlayer *RoundAction
			err := json.Unmarshal(action, &actionPlayer)
			if err != nil {
				log.Println("Error marshaling game state:", err)
				return
			}
			playersInGame := 0
			playersFinished := 0
			for client := range r.clients {
				if client.isReady == true {
					playersInGame++
				}
				if client.isReady == true && actionPlayer.Answer >= 0 {
					playersFinished++
				}
				if client.name == actionPlayer.Name {
					client.round = actionPlayer.Round
					client.answer = actionPlayer.Answer
					if actionPlayer.Answer == r.body[r.round-1].CorrectAnswer {
						client.points = client.points + 10
					}
				}
				if playersFinished >= playersInGame && playersInGame > 0 {
					r.round++
					client.round++
				}

			}
			r.SendGameState()

		}

	}
}

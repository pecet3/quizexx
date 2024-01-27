package ws

import (
	"encoding/json"
	"log"
)

type room struct {
	name     string
	category string
	clients  map[*client]bool

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
			correctAnswer: 2,
		},
		{
			question:      "test1",
			answers:       []string{"a", "b", "c", "d"},
			correctAnswer: 2,
		},
		{
			question:      "test1",
			answers:       []string{"a", "b", "c", "d"},
			correctAnswer: 2,
		},
		{
			question:      "test1",
			answers:       []string{"a", "b", "c", "d"},
			correctAnswer: 2,
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
		case msg := <-r.forward:
			for client := range r.clients {
				client.receive <- msg
			}
		case client := <-r.ready:
			client.isReady = true
			r.play = true
			log.Println("ready is ", client.name)
			r.SendGameState()

		case action := <-r.receiveAnswer:
			var actionPlayer *Player
			err := json.Unmarshal(action, &actionPlayer)
			if err != nil {
				log.Println("Error marshaling game state:", err)
				return
			}
			playersInGame := 0
			playersFinished := 0
			log.Println("map", playersFinished, playersInGame)
			for client := range r.clients {
				log.Println(";ready: ", client.isReady)
				if client.isReady == true {
					playersInGame++
				}
				if client.isReady == true && actionPlayer.Answer >= 0 {
					playersFinished++
				}

				log.Println("aadsfdsa", actionPlayer.Name)
				if client.name == actionPlayer.Name {
					log.Println("Matched")
					client.round = actionPlayer.Round
					client.answer = actionPlayer.Answer
					if actionPlayer.Answer == r.body[r.round-1].correctAnswer {
						client.points = client.points + 10
					}
				}
				if playersFinished >= playersInGame && playersInGame > 0 {
					log.Println("Finish a round, won: ", client.name)
					r.round++
					client.round++
				}
				log.Println("map", playersFinished, playersInGame)

			}
			r.SendGameState()

		}

	}
}

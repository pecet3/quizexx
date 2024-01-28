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
	game          *Game
}

func NewRoom(name string) *room {
	r := &room{
		name:          name,
		clients:       make(map[*client]bool),
		join:          make(chan *client),
		leave:         make(chan *client),
		ready:         make(chan *client),
		forward:       make(chan []byte),
		receiveAnswer: make(chan []byte),
		game:          &Game{},
	}
	log.Println(r.game)
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
func (r *room) CheckIfEveryoneIsReady() bool {
	for c := range r.clients {
		if c.isReady == false {
			return false
		}
	}
	return true
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
			log.Println("Ready")

			if ok := r.CheckIfEveryoneIsReady(); ok {
				r.game = &Game{}
				game := r.CreateGame()
				log.Println(game)
				game.SendGameState(r)

			}

		case action := <-r.receiveAnswer:
			log.Println("test")

			var actionParsed *RoundAction
			err := json.Unmarshal(action, &actionParsed)
			if err != nil {
				log.Println("Error marshaling game state:", err)
				return
			}
			playersInGame := 0
			playersFinished := 0
			for client := range r.clients {
				if client.isReady == false {
					return
				}
				r.game.State.Actions = append(r.game.State.Actions, *actionParsed)
				log.Println("added to actions: ", actionParsed)

				r.game.State.ActionsHistory = append(r.game.State.ActionsHistory, *actionParsed)

				log.Println("added to history: ", actionParsed)
				if client.isReady == true {
					playersInGame++
				}
				if client.isReady == true && actionParsed.Answer >= 0 {
					playersFinished++
				}
				log.Println(r.game.State.Round, "round")
				if client.name == actionParsed.Name {
					client.round = actionParsed.Round
					client.answer = actionParsed.Answer
					if actionParsed.Answer == r.game.Content[r.game.State.Round].CorrectAnswer {
						client.points = client.points + 10
					}
				}
				if playersFinished >= playersInGame && playersInGame > 0 {
					r.game.State.Round++
					client.round++
					log.Println("Finished the round")
				}

			}
			// r.game.SendGameState()

		}
	}
}

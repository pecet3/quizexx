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

type RoomClient struct {
	Name    string `json:"name"`
	IsReady bool   `json:"isReady"`
}

type RoomMsg struct {
	Message string       `json:"message"`
	Clients []RoomClient `json:"clients"`
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

	roomMsgBytes, err := json.Marshal(roomMsg)
	if err != nil {
		log.Println("Error marshaling game state:", err)
		return err
	}
	event := Event{
		Type:    "room_message",
		Payload: roomMsgBytes,
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

func (r *room) Run(m *manager) {
	for {
		select {
		case msg := <-r.forward:
			for client := range r.clients {
				client.receive <- msg
			}
		case client := <-r.join:
			r.clients[client] = true
			r.SendRoomMsg(client.name + "dołączył do gry")
		case client := <-r.leave:
			close(client.receive)
			delete(r.game.Players, client)
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
				game.State = r.game.NewGameState()
				r.game = game
				log.Println(game)
				game.SendGameState(r)

			}

		case action := <-r.receiveAnswer:
			log.Println("action")

			var actionParsed *RoundAction
			err := json.Unmarshal(action, &actionParsed)
			if err != nil {
				log.Println("Error marshaling game state:", err)
				return
			}
			playersInGame := 0
			playersFinished := 0
			log.Println(playersFinished, playersInGame)
			for client := range r.game.Players {
				log.Println("action 2")
				if client.isReady == false {
					return
				}
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
					if actionParsed.Answer == r.game.Content[r.game.State.Round-1].CorrectAnswer {
						client.points = client.points + 10
						log.Println("Correct answer: ", client.name)
					}
				}
				r.game.State.Actions = append(r.game.State.Actions, *actionParsed)
				r.game.State.Score = r.game.NewScore()

				if playersFinished >= playersInGame && playersInGame > 0 {
					r.game.State.Round++
					client.round++
					r.game.NewGameState()

					log.Println("Finished the round")
				}

				r.game.SendGameState(r)
			}

		}
	}
}

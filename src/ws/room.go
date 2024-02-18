package ws

import (
	"encoding/json"
	"log"
)

type room struct {
	name    string
	clients map[*client]string
	join    chan *client
	ready   chan *client
	leave   chan *client

	forward       chan []byte
	receiveAnswer chan []byte
	game          *Game
}

type RoomClient struct {
	Name    string `json:"name"`
	IsReady bool   `json:"isReady"`
}

type RoomMsgAndInfo struct {
	Message  string       `json:"message"`
	Clients  []RoomClient `json:"clients"`
	Category string       `json:"category"`
}

type SettingsGPT struct {
	name         string
	gameCategory string
	difficulty   string
	maxRounds    string
}

func NewRoom(settings SettingsGPT) *room {
	r := &room{
		name:          settings.name,
		clients:       make(map[*client]string),
		join:          make(chan *client),
		leave:         make(chan *client),
		ready:         make(chan *client),
		forward:       make(chan []byte),
		receiveAnswer: make(chan []byte),
		game:          &Game{},
	}
	return r
}

func (m *Manager) GetRoom(name string) *room {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.rooms[name]
}

func (m *Manager) CreateRoom(settings SettingsGPT) *room {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if existingRoom, ok := m.rooms[settings.name]; ok {
		return existingRoom
	}

	newRoom := NewRoom(settings)
	newRoom.game = newRoom.CreateGame(settings)
	m.rooms[settings.name] = newRoom
	log.Println("Created a room with name:")
	return newRoom
}

func (m *Manager) RemoveRoom(name string) {
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
		if !c.isReady {
			return false
		}
	}
	return true
}

func (r *room) Run(m *Manager) {
	for {
		select {
		case msg := <-r.forward:
			for client := range r.clients {
				client.receive <- msg
			}
		case client := <-r.join:
			r.clients[client] = client.name
			r.SendMsgAndInfo(client.name + " dołączył do gry")

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
			r.SendMsgAndInfo(client.name + " jest gotowy")
			if ok := r.CheckIfEveryoneIsReady(); ok {
				r.game.State = r.game.NewGameState()
				r.game.IsGame = true
				r.game.SendGameState()
			}

		case action := <-r.receiveAnswer:

			if !r.game.IsGame {
				return
			}

			var actionParsed *RoundAction
			if err := json.Unmarshal(action, &actionParsed); err != nil {
				log.Println("Error marshaling game state:", err)
				return
			}

			for client := range r.game.Players {

				client.addPoints(*actionParsed)
				r.game.State.Actions = append(r.game.State.Actions, *actionParsed)
				r.game.State.Score = r.game.NewScore()
			}
			r.game.CheckIfShouldBeNextRound()
			r.game.SendGameState()

		}
	}
}

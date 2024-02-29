package ws

import (
	"encoding/json"
	"log"
	"strconv"
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
	settings      Settings
}

type Settings struct {
	Name         string `json:"name"`
	GameCategory string `json:"category"`
	Difficulty   string `json:"difficulty"`
	MaxRounds    string `json:"maxRounds"`
}

func NewRoom(settings Settings) *room {
	r := &room{
		name:          settings.Name,
		clients:       make(map[*client]string),
		join:          make(chan *client),
		leave:         make(chan *client),
		ready:         make(chan *client),
		forward:       make(chan []byte),
		receiveAnswer: make(chan []byte),
		game:          &Game{},
		settings:      settings,
	}
	return r
}

func (m *Manager) GetRoom(name string) *room {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	log.Println(m.rooms[name], " name ", name)
	return m.rooms[name]
}

func (m *Manager) CreateRoom(settings Settings) *room {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if existingRoom, ok := m.rooms[settings.Name]; ok {
		return existingRoom
	}

	newRoom := NewRoom(settings)
	m.rooms[settings.Name] = newRoom

	log.Println("Created a room with name: ", settings.Name)
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

			if r.game.IsGame && client.isSpectator {
				err := r.SendServerMessage(client.name + " dołączył jako widz")
				if err != nil {
					return
				}
			} else {
				err := r.SendServerMessage(client.name + " dołączył do pokoju")
				if err != nil {
					return
				}
			}

			err := r.SendSettings()
			if err != nil {
				log.Println("run err send settings")
				return
			}
			if r.game.IsGame {
				_ = r.game.SendGameState()
			}
			eventBytes, err := MarshalEventToBytes[Settings](r.settings, "room_settings")
			if err != nil {
				return
			}
			client.receive <- eventBytes
			if !r.game.IsGame && !client.isSpectator {
				r.SendReadyStatus()
			}

		case client := <-r.leave:
			close(client.receive)
			delete(r.game.Players, client)
			delete(r.clients, client)

			if len(r.clients) == 0 {
				log.Println(client.name, " is leaving a room: ", r.name)
				m.RemoveRoom(r.name)
				return
			}

		case client := <-r.ready:
			if r.game.IsGame && client.isSpectator {
				r.SendServerMessage(client.name + "dołącza jako widz")
			}

			client.isReady = true
			r.SendServerMessage(client.name + " jest gotowy")
			r.SendReadyStatus()
			if ok := r.CheckIfEveryoneIsReady(); ok {
				err := r.SendServerMessage("⏳Tworzenie gry🎲, prosimy o cierpliwość...")
				if err != nil {
					log.Println("run err send server msg")
					return
				}
				err = r.SendSettings()
				if err != nil {
					log.Println("run err send settings")
					return
				}
				game := r.CreateGame()
				r.game = game
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
				log.Println("is answered: ", client.isAnswered, "round ", r.game.State.Round)

				err := r.SendServerMessage(client.name + " odpowiedział na pytanie.")
				if err != nil {
					return
				}
				client.addPointsAndToggleIsAnswered(*actionParsed)
				r.game.State.Actions = append(r.game.State.Actions, *actionParsed)
				r.game.State.Score = r.game.NewScore()

			}

			isNextRound := r.game.CheckIfShouldBeNextRound()

			isEndGame := r.game.CheckIfIsEndGame()
			if isEndGame {
				r.game = &Game{}
				r.game.IsGame = false
				_ = r.SendServerMessage("Koniec gry")
				return
			}
			if isNextRound {
				r.game.State.Round++
				for client := range r.game.Players {
					if client.isAnswered {
						client.isAnswered = false
					}
				}

				if !isEndGame {
					newState := r.game.NewGameState()
					r.game.State = newState
				}

				err := r.SendServerMessage("Rozpoczęła się nowa runda: " + strconv.Itoa(r.game.State.Round))
				if err != nil {
					return
				}
			}
			log.Println(r.game.State.Round, " round")
			r.game.SendGameState()

		}
	}
}

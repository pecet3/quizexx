package ws

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/pecet3/quizex/external"
)

type Room struct {
	name    string
	clients map[*Client]string
	join    chan *Client
	ready   chan *Client
	leave   chan *Client

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
	Language     string `json:"language"`
}

func (r *Room) CheckIfEveryoneIsReady() bool {
	for c := range r.clients {
		if !c.isReady {
			return false
		}
	}
	return true
}

func (r *Room) Run(m *Manager, external external.ExternalService) {
	for {
		select {
		case msg := <-r.forward:
			for Client := range r.clients {
				Client.receive <- msg
			}
		case Client := <-r.join:
			r.clients[Client] = Client.name

			if r.game.IsGame && Client.isSpectator {
				err := r.SendServerMessage(Client.name + " joins as spectator")
				if err != nil {
					return
				}
			} else {
				err := r.SendServerMessage(Client.name + " joins the game")
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
			Client.receive <- eventBytes
			if !r.game.IsGame && !Client.isSpectator {
				r.SendReadyStatus()
			}

		case Client := <-r.leave:
			close(Client.receive)
			delete(r.game.Players, Client)
			delete(r.clients, Client)

			if len(r.clients) == 0 {
				log.Println(Client.name, " is leaving a room: ", r.name)
				m.RemoveRoom(r.name)
				return
			}

		case Client := <-r.ready:
			if r.game.IsGame && Client.isSpectator {
				r.SendServerMessage(Client.name + "joins as a spectator")
			}

			Client.isReady = true
			r.SendServerMessage(Client.name + " jest gotowy")
			r.SendReadyStatus()
			if ok := r.CheckIfEveryoneIsReady(); ok {
				err := r.SendServerMessage("⏳Creating a Game🎲 <br> Please be patient... ")
				if err != nil {
					log.Println("run err send server msg")
					return
				}
				err = r.SendSettings()
				if err != nil {
					log.Println("run err send settings")
					return
				}
				game := CreateGame(r, external)
				log.Println(game.Content)

				r.game = game
				r.game.State = r.game.NewGameState(r.game.Content)
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

			for Client := range r.game.Players {
				if Client.name == actionParsed.Name {
					log.Println("is answered: ", Client.isAnswered, "round ", r.game.State.Round)
					if !Client.isAnswered {
						err := r.SendServerMessage(Client.name + " has answered")
						if err != nil {
							return
						}
					}

					Client.addPointsAndToggleIsAnswered(*actionParsed)
					r.game.State.Actions = append(r.game.State.Actions, *actionParsed)
					r.game.State.Score = r.game.NewScore()
				}
			}

			isNextRound := r.game.CheckIfShouldBeNextRound()

			isEndGame := r.game.CheckIfIsEndGame()
			if isEndGame {
				r.game.IsGame = false
				_ = r.SendServerMessage("It's finish the game")
				time.Sleep(800 * time.Millisecond)

				continue
			}
			if isNextRound {
				r.game.State.Round++
				for Client := range r.game.Players {
					if Client.isAnswered {
						Client.isAnswered = false
					}
				}

				if !isEndGame {
					newState := r.game.NewGameState(r.game.Content)
					r.game.State = newState
				}
				time.Sleep(800 * time.Millisecond)
				r.game.SendGameState()
				err := r.SendServerMessage("New round has began: " + strconv.Itoa(r.game.State.Round))
				if err != nil {
					return
				}
			}

		}
	}
}

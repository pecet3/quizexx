package quiz

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/logger"
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
	settings      dtos.Settings
	creatorID     int
	expiresAt     time.Time
	UUID          string
}

func (r *Room) CheckIfEveryoneIsReady() bool {
	for c := range r.clients {
		if !c.isReady {
			return false
		}
	}
	return true
}

func (r *Room) Run(m *Manager) {
	logger.Info("New room: ", r)
	for {
		if !r.expiresAt.Before(time.Now()) {
			break
		}
		select {
		case msg := <-r.forward:
			for Client := range r.clients {
				Client.receive <- msg
			}
		case Client := <-r.join:
			r.clients[Client] = Client.name
			if len(r.clients) == 0 {
				r.expiresAt = time.Now().Add(time.Hour * 2)
			}
			if r.game.IsGame && Client.isSpectator {
				err := r.sendServerMessage(Client.name + " joins as spectator")
				if err != nil {
					return
				}
			} else {
				err := r.sendServerMessage(Client.name + " joins the game")
				if err != nil {
					return
				}
			}

			err := r.sendSettings()
			if err != nil {
				log.Println("run err send settings")
				return
			}
			if r.game.IsGame {
				_ = r.game.sendGameState()
			}
			eventBytes, err := marshalEventToBytes[dtos.Settings](r.settings, "room_settings")
			if err != nil {
				return
			}
			Client.receive <- eventBytes
			if !r.game.IsGame && !Client.isSpectator {
				r.sendReadyStatus()
			}

		case Client := <-r.leave:
			close(Client.receive)
			delete(r.game.Players, Client)
			delete(r.clients, Client)
			r.sendServerMessage(Client.name + " is leaving the room")
			if len(r.clients) == 0 {
				log.Println("closing the room: ", r.name)
				m.removeRoom(r.name)
				return
			}

		case Client := <-r.ready:
			if r.game.IsGame && Client.isSpectator {
				r.sendServerMessage(Client.name + " joins as a spectator")
			}

			Client.isReady = true
			r.sendServerMessage(Client.name + " is ready!")
			r.sendReadyStatus()

			if ok := r.CheckIfEveryoneIsReady(); ok {
				err := r.sendServerMessage("⏳Creating a Game🎲 <br> Please be patient... ")
				if err != nil {
					log.Println("send server msg err: ", err)
					return
				}
				err = r.sendSettings()
				if err != nil {
					log.Println("send settings err: ", err)
					return
				}
				r.game, err = r.CreateGame()
				if err != nil {
					log.Println("create game err: ", err)
					return
				}
				r.game.State = r.game.NewGameState(r.game.Content)
				r.game.IsGame = true
				r.game.sendGameState()
			}

		case action := <-r.receiveAnswer:
			if !r.game.IsGame {
				log.Println("is game: false, ", r.game.IsGame)
				return
			}
			var actionParsed *RoundAction
			if err := json.Unmarshal(action, &actionParsed); err != nil {
				log.Println("Error marshaling game state:", err)
				return
			}

			for Client := range r.game.Players {
				if Client.name == actionParsed.Name {
					if Client.isSpectator {
						return
					}
					if !Client.isAnswered {
						err := r.sendServerMessage(Client.name + " has answered")
						if err != nil {
							return
						}
					}

					Client.addPointsAndToggleIsAnswered(*actionParsed, r)
					r.game.State.Actions = append(r.game.State.Actions, *actionParsed)
					r.game.State.Score = r.game.NewScore()
					r.game.SendPlayersAnswered()
				}
			}

			isNextRound := r.game.CheckIfShouldBeNextRound()

			indexCurrentContent := r.game.Content[r.game.State.Round-1]
			indexOkAnswr := indexCurrentContent.CorrectAnswer
			strOkAnswr := indexCurrentContent.Answers[indexOkAnswr]

			isEndGame := r.game.CheckIfIsEndGame()
			if isEndGame {
				err := r.game.sendGameState()
				if err != nil {
					log.Println("finish game err send game", err)
					continue
				}

				err = r.sendServerMessage("The correct answer was: " + strOkAnswr)
				if err != nil {
					log.Println("finish game err", err)
					continue
				}
				time.Sleep(1800 * time.Millisecond)
				_ = r.sendServerMessage("It's finish the game")

				for Client := range r.clients {
					log.Println("deleting client ", Client.name)
					close(Client.receive)
					// Client.conn.Close()
				}
				delete(m.rooms, r.name)
				return
			}

			if isNextRound {
				log.Println("It was round: ", r.game.State.Round)

				r.game.State.Round++

				var err error
				winnersStr := strings.Join(r.game.State.RoundWinners, ", ")
				log.Println("Round winners: ", len(r.game.State.RoundWinners))
				if len(r.game.State.RoundWinners) == 0 {
					err = r.sendServerMessage("No one wins this round")
				}
				if len(r.game.State.RoundWinners) == 1 {
					err = r.sendServerMessage("This round wins " + winnersStr)
				}
				if len(r.game.State.RoundWinners) >= 2 {
					err = r.sendServerMessage("This round win: " + winnersStr)
				}
				if err != nil {
					continue
				}

				if !isEndGame {
					newState := r.game.NewGameState(r.game.Content)
					r.game.State = newState
				}
				time.Sleep(1800 * time.Millisecond)
				r.game.sendGameState()

				err = r.sendServerMessage("The correct answer was: " + strOkAnswr)
				if err != nil {
					continue
				}

				time.Sleep(2800 * time.Millisecond)
				err = r.sendServerMessage("New round has began: " + strconv.Itoa(r.game.State.Round))
				if err != nil {
					continue
				}
				for Client := range r.game.Players {
					if Client.isAnswered {
						Client.isAnswered = false
					}
				}

			}
		}
	}
}

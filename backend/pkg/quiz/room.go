package quiz

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/logger"
)

type UUID = string
type Room struct {
	Name string
	UUID string

	clients map[UUID]*Client
	cMu     sync.RWMutex
	join    chan *Client
	ready   chan *Client
	leave   chan *Client

	forward       chan []byte
	receiveAnswer chan []byte
	game          *Game
	settings      dtos.Settings
	creatorID     int
	createdAt     time.Time
}

func (r *Room) CheckIfEveryoneIsReady() bool {
	for _, c := range r.clients {
		if !c.isReady {
			return false
		}
	}
	return true
}

func (r *Room) addClient(c *Client) {
	r.cMu.Lock()
	defer r.cMu.Unlock()
	r.clients[c.user.UUID] = c
}
func (r *Room) removeClient(c *Client) {
	r.cMu.Lock()
	defer r.cMu.Unlock()
	if _, ok := r.clients[c.user.UUID]; ok {
		// close connection
		delete(r.clients, c.user.UUID)
		delete(r.game.Players, c.user.UUID)
	}

}

func (r *Room) Run(m *Manager) {
	logger.Info(fmt.Sprintf(`Created a room: %s. creator user id: %d`, r.UUID, r.creatorID))
	ticker := time.NewTicker(time.Second * 20)
	defer func() {
		ticker.Stop()
		m.removeRoom(r.Name)
	}()
	for {
		select {
		case <-ticker.C:
			if len(r.clients) <= 0 {
				logger.Info(fmt.Sprintf(`No one is ine the room: %d. Closing...`, len(r.clients)))
				return
			}

		case msg := <-r.forward:
			for _, client := range r.clients {
				client.receive <- msg
			}
		case client := <-r.join:
			if len(r.clients) == 0 {
				ticker.Reset(time.Second * 20)
			}
			client.lastActive = time.Now()
			r.addClient(client)

			if r.game.IsGame && client.isSpectator {
				err := r.sendServerMessage(client.name + " joins as spectator")
				if err != nil {
					return
				}
			} else {
				err := r.sendServerMessage(client.name + " joins the game")
				if err != nil {
					return
				}
			}
			err := r.sendSettings()
			if err != nil {
				logger.Info("run err send settings")
				return
			}
			if r.game.IsGame {
				_ = r.game.sendGameState()
			}
			eventBytes, err := marshalEventToBytes[dtos.Settings](r.settings, "room_settings")
			if err != nil {
				return
			}
			client.receive <- eventBytes
			if !r.game.IsGame && !client.isSpectator {
				r.sendReadyStatus()
			}

		case client := <-r.leave:
			r.sendServerMessage(client.name + " is leaving the room")

		case client := <-r.ready:
			if r.game.IsGame && client.isSpectator {
				r.sendServerMessage(client.name + " joins as a spectator")
			}
			client.lastActive = time.Now()

			client.isReady = true
			r.sendServerMessage(client.name + " is ready!")
			r.sendReadyStatus()

			if ok := r.CheckIfEveryoneIsReady(); ok {
				err := r.sendServerMessage("Have a good game!")
				if err != nil {
					logger.Error("send server msg err: ", err)
					continue
				}
				err = r.sendSettings()
				if err != nil {
					logger.Error("send settings err: ", err)
					continue
				}
				r.game.State = r.game.newGameState(r.game.Content)
				r.game.IsGame = true
				r.game.sendGameState()
			}

		case action := <-r.receiveAnswer:
			if !r.game.IsGame {
				logger.Info("is game: false, ", r.game.IsGame)
				continue
			}
			var actionParsed *RoundAction
			if err := json.Unmarshal(action, &actionParsed); err != nil {
				logger.Error("Error marshaling game state:", err)
				continue
			}
			for _, client := range r.game.Players {
				if client.user.UUID == actionParsed.UUID {
					if client.isSpectator {
						continue
					}
					if !client.isAnswered {
						err := r.sendServerMessage(client.name + " just answered")
						if err != nil {
							continue
						}
					}
					if isGoodAnswer := r.game.checkAnswer(client, actionParsed); isGoodAnswer {
						client.points = client.points + 10
						r.game.State.RoundWinners = append(r.game.State.RoundWinners, client.name)
					}
					r.game.toggleClientIsAnswered(client, actionParsed)
					client.lastActive = time.Now()
					r.game.State.Actions = append(r.game.State.Actions, *actionParsed)
					r.game.State.Score = r.game.newScore()
					if err := r.game.sendPlayersAnswered(); err != nil {
						logger.Error(err)
						continue
					}
				}
			}

			isNextRound := r.game.checkIfShouldBeNextRound()
			indexCurrentContent := r.game.Content[r.game.State.Round-1]
			indexOkAnswr := indexCurrentContent.CorrectAnswer
			strOkAnswr := indexCurrentContent.Answers[indexOkAnswr]

			isEndGame := r.game.checkIfIsEndGame()
			if isEndGame {
				err := r.game.sendGameState()
				if err != nil {
					logger.Error("finish game err send game", err)
					continue
				}

				if err = r.sendServerMessage("The correct answer was: " + strOkAnswr); err != nil {
					logger.Error("finish game err", err)
					continue
				}
				time.Sleep(1800 * time.Millisecond)
				if err = r.sendServerMessage("It's finish the game"); err != nil {
					logger.Error("finish game err", err)
					continue
				}
				time.Sleep(1800 * time.Millisecond)
				winners := r.game.findWinner()
				winnersStr := strings.Join(winners, ", ")
				if len(winners) == 1 && len(winners) > 0 {
					if err = r.sendServerMessage("This game wins: " + winnersStr); err != nil {
						logger.Error("finish game err", err)
						continue
					}
				} else {
					if err = r.sendServerMessage("This game win: " + winnersStr); err != nil {
						logger.Error("finish game err", err)
						continue
					}
				}
				logger.Debug(winners)
				continue
			}

			if isNextRound {
				r.game.State.Round++

				var err error
				winnersStr := strings.Join(r.game.State.RoundWinners, ", ")
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
					logger.Error(err)
					continue
				}

				if !isEndGame {
					newState := r.game.newGameState(r.game.Content)
					r.game.State = newState
				}
				time.Sleep(1800 * time.Millisecond)
				r.game.sendGameState()

				err = r.sendServerMessage("The correct answer was: " + strOkAnswr)
				if err != nil {
					continue
				}

				time.Sleep(2000 * time.Millisecond)
				err = r.sendServerMessage("New round has began: " + strconv.Itoa(r.game.State.Round))
				if err != nil {
					continue
				}
				for _, client := range r.game.Players {
					if client.isAnswered {
						client.isAnswered = false
					}
				}

			}
		}
	}
}

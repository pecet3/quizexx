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

func (r *Room) checkIfEveryoneIsReady() bool {
	for _, c := range r.clients {
		if !c.player.isReady {
			return false
		}
	}
	return true
}

func (r *Room) addClient(c *Client) {
	r.cMu.Lock()
	defer r.cMu.Unlock()
	if p, ok := r.game.Players[c.user.UUID]; ok {
		c.player = p
	} else {
		c.player.user = c.user
		r.game.Players[c.user.UUID] = &Player{
			user: c.user,
		}
	}
	r.clients[c.user.UUID] = c
}
func (r *Room) removeClient(c *Client) {
	r.cMu.Lock()
	defer r.cMu.Unlock()
	delete(r.clients, c.user.UUID)
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
			r.addClient(client)

			if r.game.IsGame && client.player.isSpectator {
				err := r.sendServerMessage(client.user.Name + " joins as spectator")
				if err != nil {
					return
				}
			} else {
				err := r.sendServerMessage(client.user.Name + " joins the game")
				if err != nil {
					return
				}
			}
			if err := r.sendSettings(); err != nil {
				logger.Info("run err send settings")
				return
			}
			if r.game.IsGame {
				if err := r.game.sendGameState(); err != nil {
					logger.Error(err)
					continue
				}
			}
			eventBytes, err := marshalEventToBytes[dtos.Settings](r.settings, "room_settings")
			if err != nil {
				return
			}
			client.receive <- eventBytes
			if !r.game.IsGame && !client.player.isSpectator {
				r.sendReadyStatus()
			}

		case client := <-r.leave:
			r.sendServerMessage(client.user.Name + " is leaving the room")
			if !r.game.IsGame {
				if err := r.sendReadyStatus(); err != nil {
					logger.Error(err)
					continue
				}
			}
		case client := <-r.ready:
			if r.game.IsGame && client.player.isSpectator {
				if err := r.sendServerMessage(client.player.user.Name + " joins as a spectator"); err != nil {
					logger.Error(err)
					continue
				}
			}
			client.player.lastActive = time.Now()

			client.player.isReady = true
			if err := r.sendServerMessage(client.user.Name + " is ready!"); err != nil {
				logger.Error(err)
				continue
			}
			r.sendReadyStatus()

			if ok := r.checkIfEveryoneIsReady(); ok {
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
				if err := r.game.sendGameState(); err != nil {
					logger.Error(err)
					continue
				}
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
			for _, player := range r.game.Players {
				if player.user.UUID == actionParsed.UUID {
					if player.isSpectator {
						continue
					}
					if !player.isAnswered {
						err := r.sendServerMessage(player.user.Name + " just answered")
						if err != nil {
							continue
						}
					}
					if isGoodAnswer := r.game.checkAnswer(player, actionParsed); isGoodAnswer {
						player.points = player.points + 10
						r.game.State.RoundWinners = append(r.game.State.RoundWinners, player.user.Name)
					}
					r.game.toggleClientIsAnswered(player, actionParsed)
					player.lastActive = time.Now()
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
			logger.Debug(r.game.Players)
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
				err = r.sendServerMessage("The correct answer was: " + strOkAnswr)
				if err != nil {
					logger.Error(err)
					continue
				}

				time.Sleep(3000 * time.Millisecond)
				err = r.sendServerMessage("Round " + strconv.Itoa(r.game.State.Round) + " just started!")
				if err != nil {
					logger.Error(err)
					continue
				}

				for _, client := range r.game.Players {
					if client.isAnswered {
						client.isAnswered = false
					}
				}
				if err = r.game.sendGameState(); err != nil {
					logger.Error(err)
					continue
				}
			}
		}
	}
}

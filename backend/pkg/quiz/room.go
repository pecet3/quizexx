package quiz

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/logger"
)

const HEARTBEAT_DURATION = time.Second * 1

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
	receiveAnswer chan *RoundAction
	timeLeft      chan bool

	game      *Game
	creatorID int
	createdAt time.Time
}

func (r *Room) CreateGame(settings *dtos.Settings) (*Game, error) {
	logger.Info("Creating a game in room: ", r.Name)

	newGame := &Game{
		UUID:             uuid.NewString(),
		Room:             r,
		State:            &GameState{Round: 1},
		IsGame:           false,
		Players:          make(map[UUID]*Player),
		Settings:         settings,
		Content:          nil,
		SecLeftForAnswer: 30,
	}
	if err := newGame.getGameContent(settings); err != nil {
		return nil, err
	}
	r.game = newGame
	return newGame, nil
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
		if !r.game.IsGame {
			r.game.Players[c.user.UUID] = &Player{
				user: c.user,
			}
		}
	}
	r.clients[c.user.UUID] = c
}
func (r *Room) removeClient(c *Client) {
	r.cMu.Lock()
	defer r.cMu.Unlock()
	if _, ok := r.game.Players[c.user.UUID]; ok {
		if !r.game.IsGame {
			delete(r.game.Players, c.user.UUID)
		}
	}
	delete(r.clients, c.user.UUID)
}

func (r *Room) checkWaitRoom() error {
	if ok := r.checkIfEveryoneIsReady(); ok {
		err := r.sendServerMessage("Have a good game!")
		if err != nil {
			return err
		}
		err = r.sendSettings()
		if err != nil {
			return err

		}
		r.game.State = r.game.newGameState(r.game.Content)
		r.game.IsGame = true
		if err := r.game.sendGameState(); err != nil {
			return err
		}
	}
	return nil
}

func (r *Room) Run(m *Manager) {
	logger.Info(fmt.Sprintf(`Created a room: %s. creator user id: %d`, r.UUID, r.creatorID))
	ticker := time.NewTicker(time.Second * 20)
	heartBeat := time.NewTicker(HEARTBEAT_DURATION)
	defer func() {
		ticker.Stop()
		m.removeRoom(r.Name)
	}()

	go func(r *Room) {
		for {
			select {
			case msg := <-r.forward:
				for _, client := range r.clients {
					client.receive <- msg
				}
			case <-heartBeat.C:
				if !r.game.IsGame {
					continue
				}
				counter := r.game.GetSecLeftForAnswer()
				r.game.UpdateSecLeftForAnswer(counter - 1)
				r.sendTimeForAnswer(counter)
				if counter <= 0 {
					r.timeLeft <- true
				}
			}
		}
	}(r)
	for {
		select {
		case <-ticker.C:
			if len(r.clients) <= 0 {
				logger.Info(fmt.Sprintf(`No one is ine the room: %d. Closing...`, len(r.clients)))
				return
			}
		case client := <-r.join:
			if len(r.clients) == 0 {
				ticker.Reset(time.Second * 20)
			}
			r.addClient(client)

			if r.game.IsGame && client.isSpectator {
				err := r.sendServerMessage(client.user.Name + " joins as spectator")
				if err != nil {
					continue
				}
			} else {
				err := r.sendServerMessage(client.user.Name + " joins the game")
				if err != nil {
					continue

				}
			}
			if err := r.sendSettings(); err != nil {
				logger.Info("run err send settings")
				continue

			}
			if r.game.IsGame {
				if err := r.game.sendGameState(); err != nil {
					logger.Error(err)
					continue
				}
			}
			if err := r.sendSettings(); err != nil {
				continue
			}
			if !r.game.IsGame && !client.isSpectator {
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
			if err := r.checkWaitRoom(); err != nil {
				logger.Error(err)
				continue
			}
		case client := <-r.ready:
			if r.game.IsGame && client.isSpectator {
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

			if err := r.checkWaitRoom(); err != nil {
				logger.Error(err)
				continue
			}

		case action := <-r.receiveAnswer:
			if !r.game.IsGame {
				continue
			}
			if player, ok := r.game.Players[action.UUID]; ok {
				if !player.isAnswered {
					err := r.sendServerMessage(player.user.Name + " just answered")
					if err != nil {
						continue
					}
				} else {
					continue
				}
				if isGoodAnswer := r.game.checkAnswer(player, action); isGoodAnswer {
					r.game.State.RoundWinners = append(r.game.State.RoundWinners, player.user.UUID)
				}
				r.game.toggleClientIsAnswered(player, action)
				player.lastActive = time.Now()
				r.game.State.Actions = append(r.game.State.Actions, action)
				if err := r.game.sendPlayersAnswered(); err != nil {
					logger.Error(err)
					continue
				}

				if err := r.game.performRound(heartBeat, false); err != nil {
					logger.Error(err)
					continue
				}
			}

		case <-r.timeLeft:
			if !r.game.IsGame {
				continue
			}
			err := r.sendServerMessage("Time for answer left!")
			if err != nil {
				continue
			}

			if err := r.game.performRound(heartBeat, true); err != nil {
				logger.Error(err)
				continue
			}
		}
	}
}

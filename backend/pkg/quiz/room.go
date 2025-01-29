package quiz

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/pkg/logger"
)

const (
	HEARTBEAT_DURATION      = time.Second * 1
	NOBODYCHECKING_DURATION = time.Second * 10

	// time for reading
	TFR_SHORT_DURATION = time.Millisecond * 1800
	TFR_LONG_DURATION  = time.Millisecond * 2500
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
	receiveAnswer chan *RoundAction
	timeLeft      chan bool

	game      *Game
	creatorID int
	createdAt time.Time

	gcDb   *data.GameContent
	gcrsDb []*data.GameContentRound
	gDb    *data.Game
}

func (r *Room) checkWaitRoom(m *Manager) error {
	if ok := r.checkIfEveryoneIsReady(); ok {
		err := r.sendServerMessage("Have a good game!")
		if err != nil {
			return err
		}
		err = r.sendSettings()
		if err != nil {
			return err

		}
		newState := r.game.newGameState(m.d, r.game.Content)
		r.game.State = newState
		ctx := context.Background()
		gc, err := m.d.GetGameContentByUUID(ctx, r.game.UUID)
		if err != nil {
			return err
		}
		_, err = m.d.AddGame(ctx, data.AddGameParams{
			RoomUuid:      r.UUID,
			RoomName:      r.Name,
			GameContentID: gc.ID,
		})
		if err != nil {
			return err
		}
		r.game.IsGame = true
		if err := r.game.sendGameState(); err != nil {
			return err
		}
	}
	return nil
}

func (r *Room) Run(m *Manager) {
	logger.Info(fmt.Sprintf(`Created a room: %s. creator user id: %d`, r.UUID, r.creatorID))

	gDb, err := m.d.AddGame(context.Background(), data.AddGameParams{
		RoomUuid:      r.UUID,
		RoomName:      r.Name,
		GameContentID: r.gcDb.ID,
	})
	if err != nil {
		logger.Error(err)
		return
	}
	r.gDb = &gDb

	nobodyCheckingT := time.NewTicker(NOBODYCHECKING_DURATION)
	heartBeatT := time.NewTicker(HEARTBEAT_DURATION)

	logger.Debug(r.gcDb)
	logger.Debug(r.gcrsDb)
	logger.Debug(r.gcrsDb[0].Round)

	defer func() {
		nobodyCheckingT.Stop()
		heartBeatT.Stop()
		m.removeRoom(r.Name)
	}()
	go func(r *Room) {
		for {
			select {
			case msg := <-r.forward:
				for _, client := range r.clients {
					client.receive <- msg
				}
			case <-heartBeatT.C:
				if !r.game.IsGame {
					continue
				}
				counter := r.game.getSecLeftForAnswer()
				r.game.updateSecLeftForAnswer(counter - 1)
				if counter >= 0 {
					r.sendTimeForAnswer(counter)
				}
				if counter == 0 {
					r.game.updateSecLeftForAnswer(-1)
					r.timeLeft <- true
				}

			}

		}
	}(r)

	for {
		select {
		case <-nobodyCheckingT.C:
			if len(r.clients) <= 0 {
				logger.Info(fmt.Sprintf(`No one is ine the room: %d. Closing...`, len(r.clients)))
				return
			}
		case client := <-r.join:
			if len(r.clients) == 0 {
				nobodyCheckingT.Reset(time.Second * 20)
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
			if err := r.checkWaitRoom(m); err != nil {
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

			if err := r.checkWaitRoom(m); err != nil {
				logger.Error(err)
				continue
			}

		case action := <-r.receiveAnswer:
			timeLeft := r.game.getSecLeftForAnswer()
			if !r.game.IsGame || timeLeft <= 0 || action.Round != r.game.State.Round {
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
				isGoodAnswer := r.game.checkAnswer(player, action)
				if isGoodAnswer {
					r.game.State.RoundWinners = append(r.game.State.RoundWinners, player.user.Uuid)
				}
				r.game.toggleClientIsAnswered(player, action)
				player.lastActive = time.Now()
				r.game.State.Actions = append(r.game.State.Actions, action)
				if err := r.game.sendPlayersAnswered(); err != nil {
					logger.Error(err)
				}
				answer, err := m.d.GetGameContentAnswerByRoundIDAndIndex(context.Background(), data.GetGameContentAnswerByRoundIDAndIndexParams{
					GameContentRoundID: r.getDbGameConontentRoundID(),
					IndexInArr:         int64(action.Answer),
				})
				if err != nil {
					logger.Error(err)
				}
				gr, err := m.d.AddGameRoundAction(context.Background(), data.AddGameRoundActionParams{
					AnswerID:           answer.ID,
					Points:             10,
					GameID:             r.gDb.ID,
					UserID:             player.user.ID,
					AnswerIndex:        answer.IndexInArr,
					IsGoodAnswer:       isGoodAnswer,
					GameContentRoundID: answer.GameContentRoundID,
				})
				logger.Debug(gr, answer)
				if err != nil {
					logger.Error(err)
				}
				if err := r.game.PerformRound(m, heartBeatT, false); err != nil {
					logger.Error(err)
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
			time.Sleep(TFR_SHORT_DURATION)
			if err := r.game.PerformRound(m, heartBeatT, true); err != nil {
				logger.Error(err)
				continue
			}
		}
	}
}

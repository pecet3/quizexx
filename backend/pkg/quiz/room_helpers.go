package quiz

import (
	"github.com/google/uuid"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/fetchers"
	"github.com/pecet3/quizex/pkg/logger"
)

func (r *Room) SetDbStructs(gc *data.GameContent, gcrs []*data.GameContentRound) {
	r.gcDb = gc
	r.gcrsDb = gcrs
}

func (r *Room) CreateGame(f fetchers.Fetchable, settings *dtos.Settings) (*Game, error) {
	logger.Info("Creating a game in room: ", r.Name)
	newGame := &Game{
		UUID:             uuid.NewString(),
		Room:             r,
		State:            &GameState{Round: 1},
		IsGame:           false,
		Players:          make(map[UUID]*Player),
		Settings:         settings,
		Content:          nil,
		SecLeftForAnswer: settings.SecForAnswer,
	}
	if err := newGame.getGameContent(f, settings); err != nil {
		return nil, err
	}
	r.game = newGame
	return newGame, nil
}

func (r *Room) checkIfEveryoneIsReady() bool {
	if len(r.clients) <= 0 {
		return false
	}
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
	if p, ok := r.game.Players[c.user.Uuid]; ok {
		c.player = p
	} else {
		c.player.user = c.user
		if !r.game.IsGame {
			r.game.Players[c.user.Uuid] = &Player{
				user: c.user,
			}
		}
	}
	r.clients[c.user.Uuid] = c
}
func (r *Room) removeClient(c *Client) {
	r.cMu.Lock()
	defer r.cMu.Unlock()
	if _, ok := r.game.Players[c.user.Uuid]; ok {
		if !r.game.IsGame {
			delete(r.game.Players, c.user.Uuid)
		}
	}
	delete(r.clients, c.user.Uuid)
}

func (r *Room) getDbGameConontentRoundID() int64 {
	return r.gcrsDb[r.game.State.Round-1].ID
}

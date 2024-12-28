package quiz

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pecet3/quizex/data/entities"
)

var (
	pongWait     = 60 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type Client struct {
	user        *entities.User
	name        string
	conn        *websocket.Conn
	receive     chan []byte
	room        *Room
	isReady     bool
	isSpectator bool
	isAnswered  bool
	answer      int
	points      int
	roundsWon   []uint
}

func (c *Client) addPointsAndToggleIsAnswered(action RoundAction, r *Room) {
	if c.name == action.Name {
		c.answer = action.Answer
		if action.Answer == c.room.game.Content[c.room.game.State.Round-1].CorrectAnswer && !c.isAnswered {
			c.points = c.points + 10
			r.game.State.RoundWinners = append(r.game.State.RoundWinners, c.name)
		}
		if action.Answer >= 0 && !c.isAnswered {
			c.room.game.State.PlayersAnswered = append(c.room.game.State.PlayersAnswered, c.name)
			c.isAnswered = true
		}
	}
}

func (c *Client) read() {
	defer func() {
	}()

	if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
		return
	}

	c.conn.SetReadLimit(1024)

	c.conn.SetPongHandler(c.pongHandler)

	for {
		_, reqBytes, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("ws err or user was too long inactive:", err)
				c.room.leave <- c
				return
			}
			log.Println("leave:", err)
			c.room.leave <- c
			return
		}

		var request Event
		if err := json.Unmarshal(reqBytes, &request); err != nil {
			log.Println("error marshaling json", err)
			break
		}
		if request.Type == "ready_player" {
			c.room.ready <- c
		}
		if request.Type == "send_answer" {
			var actionPlayer *RoundAction
			err := json.Unmarshal(request.Payload, &actionPlayer)
			if err != nil {
				log.Println("Error marshaling game state:", err)
				return
			}
			c.room.receiveAnswer <- request.Payload
		}
		if request.Type == "chat_message" {
			c.room.forward <- reqBytes
		}
	}
}

func (c *Client) write() {
	defer func() {
		c.conn.Close()
	}()
	ticker := time.NewTicker(pingInterval)

	for {
		select {
		case msg, ok := <-c.receive:
			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed: ", err)
				}
				return

			}
			err := c.conn.WriteMessage(websocket.TextMessage, msg)
			if err == websocket.ErrCloseSent {
				log.Println(err)
				continue
			}
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				return
			}

		case <-ticker.C:
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte(``)); err != nil {
				log.Println("write message error: ", err)
				return
			}
		}
	}
}

func (c *Client) pongHandler(pongMsg string) error {
	return c.conn.SetReadDeadline(time.Now().Add(pongWait))
}

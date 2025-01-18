package quiz

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pecet3/quizex/data/entities"
	"github.com/pecet3/quizex/pkg/logger"
)

var (
	pongWait     = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type Client struct {
	user    *entities.User
	conn    *websocket.Conn
	receive chan []byte
	room    *Room
	player  *Player
}

func (c *Client) read(r *Room) {
	defer func() {
		r.removeClient(c)
	}()

	if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		logger.Error(err)
		return
	}

	c.conn.SetReadLimit(4096)

	c.conn.SetPongHandler(c.pongHandler)

	for {
		_, reqBytes, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Error("ws err or user was too long inactive:", err)
				return
			}
			logger.Error(err)
			return
		}

		var request Event
		if err := json.Unmarshal(reqBytes, &request); err != nil {
			logger.Error("error marshaling json", err)
			continue
		}
		if request.Type == "ready_player" {
			c.room.ready <- c
		}
		if request.Type == "send_answer" {
			var actionPlayer *RoundAction
			err := json.Unmarshal(request.Payload, &actionPlayer)
			if err != nil {
				logger.Error("Error marshaling game state:", err)
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
					logger.Error("connection closed: ", err)
				}
				return
			}
			err := c.conn.WriteMessage(websocket.TextMessage, msg)
			if err == websocket.ErrCloseSent {
				logger.Error(err)
				continue
			}
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				return
			}

		case <-ticker.C:
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte(``)); err != nil {
				logger.Error("write message error: ", err)
				return
			}
		}
	}
}

func (c *Client) pongHandler(pongMsg string) error {
	return c.conn.SetReadDeadline(time.Now().Add(pongWait))
}

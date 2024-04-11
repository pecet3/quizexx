package ws

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pongWait     = 60 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type Client struct {
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

func (Client *Client) addPointsAndToggleIsAnswered(action RoundAction) {
	if Client.name == action.Name {
		Client.answer = action.Answer
		if action.Answer == Client.room.game.Content[Client.room.game.State.Round-1].CorrectAnswer && !Client.isAnswered {
			Client.points = Client.points + 10
		}
		if action.Answer >= 0 && !Client.isAnswered {
			Client.room.game.State.PlayersAnswered = append(Client.room.game.State.PlayersAnswered, Client.name)
			Client.isAnswered = true
		}
	}
}

func (c *Client) read() {
	defer func() {
		c.conn.Close()
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
			log.Println("ws err !!!!:", err)

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
		case msg := <-c.receive:

			err := c.conn.WriteMessage(websocket.TextMessage, msg)

			if err != nil {
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

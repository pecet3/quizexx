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

type client struct {
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

func (client *client) addPointsAndToggleIsAnswered(action RoundAction) {
	if client.name == action.Name {
		client.answer = action.Answer
		if action.Answer == client.room.game.Content[client.room.game.State.Round-1].CorrectAnswer && !client.isAnswered {
			client.points = client.points + 10
		}
		if action.Answer >= 0 && !client.isAnswered {
			client.room.game.State.PlayersFinished = append(client.room.game.State.PlayersFinished, client.name)
			client.isAnswered = true
		}
	}
}

func (c *client) read(m *Manager) {
	defer func() {
		c.conn.Close()
	}()

	if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {

		log.Println(err)
		return
	}

	c.conn.SetReadLimit(512)

	c.conn.SetPongHandler(c.pongHandler)

	for {
		_, payload, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("ws err or user was too long inactive:", err)
				c.room.leave <- c
				return
			}
			log.Println("ws err !!!!:", err)

		}

		var request Event
		if err := json.Unmarshal(payload, &request); err != nil {
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
	}
}

func (c *client) write() {
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

func (c *client) pongHandler(pongMsg string) error {
	return c.conn.SetReadDeadline(time.Now().Add(pongWait))
}

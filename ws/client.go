package ws

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pongWait     = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type client struct {
	name    string
	conn    *websocket.Conn
	receive chan []byte
	room    *room
	isReady bool
	round   int
	answer  int
	points  int
}

func (c *client) read(m *manager) {
	defer c.conn.Close()

	c.conn.SetReadLimit(512)

	c.conn.SetPongHandler(c.pongHandler)

	for {
		_, payload, err := c.conn.ReadMessage()
		if err != nil {
			return
		}
		var request Event
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Println("error marshaling json", err)
			continue
		}
		if request.Type == "ready_player" {
			c.room.ready <- c
		}
		if request.Type == "send_answer" {
			log.Println(request.Payload)
			c.room.receiveAnswer <- request.Payload
		}
	}
}

func (c *client) write() {
	defer c.conn.Close()

	for msg := range c.receive {
		log.Println(msg, "aaaa")

		log.Println("new message in room: ", c.room.name)

		err := c.conn.WriteMessage(websocket.TextMessage, msg)

		if err != nil {
			return
		}
	}
}

func (c *client) pongHandler(pongMsg string) error {
	log.Println("Pong")
	return c.conn.SetReadDeadline(time.Now().Add(pongWait))
}

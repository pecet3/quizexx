package ws

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
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

	for {
		_, payload, err := c.conn.ReadMessage()
		if err != nil {
			return
		}
		var request Event
		log.Println(string(payload))
		log.Println(request)
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Println("error marshaling json", err)
			continue
		}
		log.Println(request.Payload)
		if request.Type == "ready_player" {
			c.room.ready <- c
		}
		if request.Type == "send_answer" {
			c.room.receiveAnswer <- request.Payload
		}
		c.room.forward <- request.Payload
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

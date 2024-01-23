package ws

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/pecet3/czatex/utils"
)

type client struct {
	name    string
	conn    *websocket.Conn
	receive chan []byte
	room    *room
	isReady bool
	round   chan int
	answer  chan int
}

func (c *client) read(m *manager) {
	defer c.conn.Close()

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

		log.Println(request)
		c.room.forward <- request.Payload
	}
}

func (c *client) write() {
	defer c.conn.Close()

	for msg := range c.receive {

		result, err := utils.DecodeMessage(msg)

		var wg sync.WaitGroup
		namesChan := make(chan []string)

		wg.Add(1)

		go createNamesArr(c.room.clients, &wg, namesChan)

		namesArr := <-namesChan
		close(namesChan)

		message, err := utils.MarshalJsonMessage(result.Name, result.Message, namesArr)

		log.Println("new message in room: ", c.room.name)

		err = c.conn.WriteMessage(websocket.TextMessage, message)

		if err != nil {
			return
		}
	}
}

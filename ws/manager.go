package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type manager struct {
	rooms  map[string]*room
	mutex  sync.Mutex
	events map[string]EventHandler
}

var (
	upgrader = &websocket.Upgrader{
		CheckOrigin:     checkOrigin,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func checkOrigin(r *http.Request) bool {
	return true
}

func NewManager() *manager {
	m := &manager{
		rooms: make(map[string]*room),
	}
	return m
}

func (m *manager) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
		return
	}

	room := req.URL.Query().Get("room")
	if room == "" {
		return
	}
	name := req.URL.Query().Get("name")
	if name == "" || name == "serwer" || name == "klient" {
		return
	}

	log.Printf("New connection: %v connected to room: %v", name, room)

	currentRoom := m.GetRoom(room)

	if currentRoom == nil {
		currentRoom = m.CreateRoom(room)
		go currentRoom.Run(m)
	}

	client := &client{
		conn:    conn,
		receive: make(chan []byte),
		room:    currentRoom,
		name:    name,
		round:   0,
		answer:  -1,
		points:  0,
		isReady: false,
	}

	currentRoom.join <- client
	defer func() { currentRoom.leave <- client }()
	go client.write()
	client.read(m)
}

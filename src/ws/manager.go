package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Manager struct {
	mutex sync.Mutex
	rooms map[string]*Room
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

func NewManager() *Manager {
	return &Manager{
		rooms: make(map[string]*Room),
	}
}

func (m *Manager) GetRoomNamesList() []string {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	var names []string

	for roomName := range m.rooms {
		names = append(names, roomName)
	}

	return names
}

func (m *Manager) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
		return
	}

	roomName := req.URL.Query().Get("room")
	if roomName == "" {
		return
	}
	name := req.URL.Query().Get("name")
	if name == "" || name == "serwer" || name == "klient" {
		return
	}
	difficulty := req.URL.Query().Get("difficulty")
	maxRounds := req.URL.Query().Get("maxRounds")
	category := req.URL.Query().Get("category")
	newRoom := req.URL.Query().Get("new")

	if len(category) >= 32 || len(category) < 5 {
		conn.Close()
		return
	}

	settings := Settings{
		Name:         roomName,
		GameCategory: category,
		Difficulty:   difficulty,
		MaxRounds:    maxRounds,
	}
	currentRoom := m.GetRoom(roomName)

	if currentRoom != nil {
		if newRoom == "true" {
			conn.Close()
			return
		}

		for roomClient := range currentRoom.clients {
			if name == roomClient.name {
				conn.Close()
				break
			}
		}
	}

	if currentRoom == nil {
		if newRoom == "true" {
			currentRoom = m.CreateRoom(settings)
			go currentRoom.Run(m)
		} else {
			conn.Close()
			return
		}
	}

	isSpectator := false

	if currentRoom.game.IsGame {
		isSpectator = true
	}

	client := &client{
		conn:        conn,
		receive:     make(chan []byte),
		room:        currentRoom,
		name:        name,
		answer:      -1,
		points:      0,
		isReady:     false,
		isSpectator: isSpectator,
		isAnswered:  false,
	}

	log.Printf("New connection, userName: %v connected to room: %v", name, roomName)
	currentRoom.join <- client
	defer func() { currentRoom.leave <- client }()
	go client.write()
	client.read(m)
}

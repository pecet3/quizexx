package quiz

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/external"
)

type Manager struct {
	mu       sync.Mutex
	ctx      context.Context
	rooms    map[string]*Room
	external *external.ExternalService
}

func NewManager() *Manager {
	return &Manager{
		rooms:    make(map[string]*Room),
		mu:       sync.Mutex{},
		external: &external.ExternalService{},
	}
}

func (m *Manager) newRoom(settings dtos.Settings, creatorID int) *Room {
	r := &Room{
		name:          settings.Name,
		clients:       make(map[*Client]string),
		join:          make(chan *Client),
		leave:         make(chan *Client),
		ready:         make(chan *Client),
		forward:       make(chan []byte),
		receiveAnswer: make(chan []byte),
		game:          &Game{},
		settings:      settings,
		creatorID:     creatorID,
	}
	return r
}

func (m *Manager) getRoom(name string) *Room {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.rooms[name]
}

func (m *Manager) CreateRoom(settings dtos.Settings, creatorID int) *Room {
	m.mu.Lock()
	defer m.mu.Unlock()

	if existingRoom, ok := m.rooms[settings.Name]; ok {
		return existingRoom
	}

	newRoom := m.newRoom(settings, creatorID)
	m.rooms[settings.Name] = newRoom

	log.Println("> Created a room with name: ", settings.Name)
	return newRoom
}

func (m *Manager) removeRoom(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if room, ok := m.rooms[name]; ok {
		for Client := range room.clients {
			room.leave <- Client
		}
		close(room.join)
		close(room.forward)
		close(room.ready)
		close(room.receiveAnswer)
		close(room.leave)

		delete(m.rooms, name)
		log.Println("> Closing a room with name:", room.name)
		return
	}
}
func (m *Manager) GetRoomNamesList() []string {
	m.mu.Lock()
	defer m.mu.Unlock()

	var names []string

	for roomName := range m.rooms {
		names = append(names, roomName)
	}

	return names
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
func (m *Manager) ServeWs(w http.ResponseWriter, req *http.Request) {
	m.ctx = req.Context()

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
	lang := req.URL.Query().Get("lang")

	settings := dtos.Settings{
		Name:       roomName,
		GenContent: category,
		Difficulty: difficulty,
		MaxRounds:  maxRounds,
		Language:   lang,
	}
	currentRoom := m.getRoom(roomName)

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
	// to fix
	if currentRoom == nil {
		if newRoom == "true" {
			currentRoom = m.CreateRoom(settings, 0)
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

	client := &Client{
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

	log.Printf("> UserName: %v connected to room: %v", name, roomName)
	currentRoom.join <- client
	defer func() { currentRoom.leave <- client }()
	go client.write()
	client.read()
}
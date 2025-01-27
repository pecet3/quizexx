package quiz

import (
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/external"
	"github.com/pecet3/quizex/pkg/logger"
	"github.com/pecet3/quizex/pkg/social"
)

type Manager struct {
	mu       sync.Mutex
	rooms    map[string]*Room
	external *external.ExternalService
	d        *data.Queries
	s        *social.Social
}

func NewManager(d *data.Queries, s *social.Social) *Manager {
	m := &Manager{
		rooms:    make(map[string]*Room),
		mu:       sync.Mutex{},
		external: &external.ExternalService{},
		d:        d,
		s:        s,
	}

	return m
}

func (m *Manager) newRoom(name string, creatorID int) *Room {
	r := &Room{
		clients: make(map[UUID]*Client),
		join:    make(chan *Client),
		leave:   make(chan *Client),
		ready:   make(chan *Client),

		forward:       make(chan []byte),
		receiveAnswer: make(chan *RoundAction),
		timeLeft:      make(chan bool),

		game:      &Game{},
		creatorID: creatorID,
		createdAt: time.Now(),
		UUID:      uuid.NewString(),
		Name:      name,
	}
	return r
}

func (m *Manager) GetRoom(name string) *Room {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.rooms[name]
}
func (m *Manager) getRoomByUserID(uID int) *Room {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, room := range m.rooms {
		if room.creatorID == uID {
			return room
		}
	}
	return nil
}
func (m *Manager) CheckUserHasRoom(uID int) bool {
	r := m.getRoomByUserID(uID)
	return r != nil
}
func (m *Manager) CreateRoom(name string, creatorID int) *Room {
	m.mu.Lock()
	defer m.mu.Unlock()

	if existingRoom, ok := m.rooms[name]; ok {
		return existingRoom
	}

	newRoom := m.newRoom(name, creatorID)
	m.rooms[name] = newRoom

	logger.Info("Created a room with name: ", name)
	return newRoom
}

func (m *Manager) removeRoom(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if room, ok := m.rooms[name]; ok {
		for _, c := range room.clients {
			room.removeClient(c)
		}
		delete(m.rooms, name)
	}
	logger.Debug(m.rooms)
}
func (m *Manager) GetRoomsList() []*dtos.Room {
	m.mu.Lock()
	defer m.mu.Unlock()

	var rooms []*dtos.Room

	for uuid, room := range m.rooms {
		r := &dtos.Room{
			UUID:       uuid,
			Name:       room.Name,
			Players:    len(room.clients),
			MaxPlayers: 32,
			Round:      room.game.State.Round,
			MaxRounds:  room.game.Settings.MaxRounds,
		}
		rooms = append(rooms, r)
	}

	return rooms
}

var (
	upgrader = &websocket.Upgrader{
		CheckOrigin:     checkOrigin,
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
	}
)

func checkOrigin(r *http.Request) bool {
	return true
}
func (m *Manager) ServeQuiz(w http.ResponseWriter, req *http.Request, u *data.User) {
	roomName := req.PathValue("name")
	currentRoom := m.GetRoom(roomName)
	logger.Debug("room name: ", roomName)
	if currentRoom == nil {
		logger.Error("no room with provided uuid")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	isSpectator := false
	if _, ok := currentRoom.game.Players[u.Uuid]; !ok && currentRoom.game.IsGame {
		isSpectator = true
	}
	// to do if client exists,change only conn

	client := &Client{
		conn:        conn,
		receive:     make(chan []byte),
		room:        currentRoom,
		player:      &Player{},
		isSpectator: isSpectator,
		user:        u,
	}

	currentRoom.join <- client
	defer func() {
		currentRoom.leave <- client
	}()
	go client.write()
	client.read(currentRoom)

}

package quiz

import (
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pecet3/quizex/data"
	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/data/entities"
	"github.com/pecet3/quizex/pkg/external"
	"github.com/pecet3/quizex/pkg/logger"
)

type Manager struct {
	mu       sync.Mutex
	rooms    map[string]*Room
	external *external.ExternalService
	d        *data.Data
}

func NewManager(d *data.Data) *Manager {
	return &Manager{
		rooms:    make(map[string]*Room),
		mu:       sync.Mutex{},
		external: &external.ExternalService{},
		d:        d,
	}
}

func (m *Manager) newRoom(settings *dtos.Settings, creatorID int) *Room {
	r := &Room{
		Name:          settings.Name,
		clients:       make(map[UUID]*Client),
		join:          make(chan *Client),
		leave:         make(chan *Client),
		ready:         make(chan *Client),
		forward:       make(chan []byte),
		receiveAnswer: make(chan []byte),
		game:          &Game{},
		settings:      *settings,
		creatorID:     creatorID,
		createdAt:     time.Now(),
		UUID:          uuid.NewString(),
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
func (m *Manager) CreateRoom(settings *dtos.Settings, creatorID int) *Room {
	m.mu.Lock()
	defer m.mu.Unlock()

	if existingRoom, ok := m.rooms[settings.Name]; ok {
		return existingRoom
	}

	newRoom := m.newRoom(settings, creatorID)
	m.rooms[newRoom.Name] = newRoom

	logger.Info("Created a room with name: ", settings.Name)
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
		return
	}
}
func (m *Manager) GetRoomsList() []*dtos.Room {
	m.mu.Lock()
	defer m.mu.Unlock()

	var rooms []*dtos.Room

	for uuid, room := range m.rooms {
		r := &dtos.Room{
			UUID:       uuid,
			Name:       room.Name,
			Players:    len(room.game.Players),
			MaxPlayers: 10,
			Round:      room.game.State.Round,
			MaxRounds:  room.game.MaxRounds,
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
func (m *Manager) ServeQuiz(w http.ResponseWriter, req *http.Request, u *entities.User) {
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

	// isSpectator := false
	// if currentRoom.game.IsGame {
	// 	isSpectator = true
	// }
	// to do if client exists,change only conn

	client := &Client{
		conn:    conn,
		receive: make(chan []byte),
		room:    currentRoom,
		player:  &Player{},
		user:    u,
	}

	currentRoom.join <- client
	defer func() {
		currentRoom.leave <- client
	}()
	go client.write()
	client.read(currentRoom)

}

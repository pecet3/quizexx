package ws

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Manager struct {
	mutex sync.Mutex
	rooms map[string]*Room
}

func (m *Manager) NewManager() *Manager {
	return &Manager{
		rooms: make(map[string]*Room),
		mutex: sync.Mutex{},
	}
}

type IManager interface {
	HandleWs(conn *websocket.Conn, settings Settings, newRoom string, userName string)

	CreateRoom(settings Settings) *Room
	RemoveRoom(name string)
	GetRoomNamesList() []string
	GetRoom(name string) *Room
	NewRoom(settings Settings) *Room
}

func (m *Manager) NewRoom(settings Settings) *Room {
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
	}
	return r
}

func (m *Manager) GetRoom(name string) *Room {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	log.Println(m.rooms[name], " name ", name)
	return m.rooms[name]
}

func (m *Manager) CreateRoom(settings Settings) *Room {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if existingRoom, ok := m.rooms[settings.Name]; ok {
		return existingRoom
	}

	newRoom := m.NewRoom(settings)
	m.rooms[settings.Name] = newRoom

	log.Println("Created a room with name: ", settings.Name)
	return newRoom
}

func (m *Manager) RemoveRoom(name string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

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
		log.Println("Closing a room with name:", room.name)
		return
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

// func (m *Manager) ServeHTTP(w http.ResponseWriter, req *http.Request) {

// 	conn, err := upgrader.Upgrade(w, req, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	roomName := req.URL.Query().Get("room")
// 	if roomName == "" {
// 		return
// 	}
// 	name := req.URL.Query().Get("name")
// 	if name == "" || name == "serwer" || name == "klient" {
// 		return
// 	}
// 	difficulty := req.URL.Query().Get("difficulty")
// 	maxRounds := req.URL.Query().Get("maxRounds")
// 	category := req.URL.Query().Get("category")
// 	newRoom := req.URL.Query().Get("new")

// 	if len(category) >= 32 || len(category) < 5 {
// 		conn.Close()
// 		return
// 	}

// 	settings := Settings{
// 		Name:         roomName,
// 		GameCategory: category,
// 		Difficulty:   difficulty,
// 		MaxRounds:    maxRounds,
// 	}
// 	currentRoom := m.GetRoom(m, roomName)

// 	if currentRoom != nil {
// 		if newRoom == "true" {
// 			conn.Close()
// 			return
// 		}

// 		for roomClient := range currentRoom.clients {
// 			if name == roomClient.name {
// 				conn.Close()
// 				break
// 			}
// 		}
// 	}

// 	if currentRoom == nil {
// 		if newRoom == "true" {
// 			currentRoom = m.CreateRoom(settings)
// 			go currentRoom.Run(m)
// 		} else {
// 			conn.Close()
// 			return
// 		}
// 	}

// 	isSpectator := false

// 	if currentRoom.game.IsGame {
// 		isSpectator = true
// 	}

// 	client := &Client{
// 		conn:        conn,
// 		receive:     make(chan []byte),
// 		room:        currentRoom,
// 		name:        name,
// 		answer:      -1,
// 		points:      0,
// 		isReady:     false,
// 		isSpectator: isSpectator,
// 		isAnswered:  false,
// 	}

// 	log.Printf("New connection, userName: %v connected to room: %v", name, roomName)
// 	currentRoom.join <- client
// 	defer func() { currentRoom.leave <- client }()
// 	go client.write()
// 	client.read()
// }

func (m *Manager) HandleWs(conn *websocket.Conn, settings Settings, newRoom string, userName string) {
	currentRoom := m.GetRoom(settings.Name)

	if currentRoom != nil {
		if newRoom == "true" {
			conn.Close()
			return
		}

		for roomClient := range currentRoom.clients {
			if userName == roomClient.name {
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

	client := &Client{
		conn:        conn,
		receive:     make(chan []byte),
		room:        currentRoom,
		name:        userName,
		answer:      -1,
		points:      0,
		isReady:     false,
		isSpectator: isSpectator,
		isAnswered:  false,
	}
	log.Println("new conn")
	currentRoom.join <- client
	defer func() { currentRoom.leave <- client }()
	go client.write()
	client.read()
}

package ws

import (
	"encoding/json"
	"log"
	"sync"
)

type Game struct {
	Room     *room
	State    *GameState
	Content  []QandA
	Category string
	IsGame   bool
	Players  map[*client]string
	mutex    sync.Mutex
}

type GameState struct {
	Round           int           `json:"round"`
	Question        string        `json:"question"`
	Answers         []string      `json:"answers"`
	Actions         []RoundAction `json:"actions"`
	Score           []PlayerScore `json:"score"`
	PlayersFinished []string      `json:"playersFinished"`
}

type RoundAction struct {
	Name   string `json:"name"`
	Answer int    `json:"answer"`
	Round  int    `json:"round"`
}

type PlayerScore struct {
	Name      string `json:"name"`
	Points    int    `json:"points"`
	RoundsWon []uint `json:"roundsWon"`
}

type QandA struct {
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
	CorrectAnswer int      `json:"correctAnswer"`
}

func (r *room) CreateGame() *Game {
	log.Println("creating a game")

	content := []QandA{
		{
			Question:      "Co oznacza skrót CPU?",
			Answers:       []string{"Centralna Jednostka Przetwarzania", "Komputerowa Jednostka Przetwarzania", "Centralna Jednostka Procesora", "Komputerowa Jednostka Procesora"},
			CorrectAnswer: 0,
		},
		{
			Question:      "Jaki jest główny cel systemu operacyjnego?",
			Answers:       []string{"Zarządzanie zasobami sprzętowymi", "Uruchamianie aplikacji", "Zakładanie łączności internetowej", "Przechowywanie danych"},
			CorrectAnswer: 0,
		},
		{
			Question:      "Co oznacza skrót HTML?",
			Answers:       []string{"HyperText Markup Language", "HyperText Modeling Language", "High-Level Text Language", "Hyperlink and Text Markup Language"},
			CorrectAnswer: 0,
		},
		{
			Question:      "Jaka jest binarna reprezentacja liczby dziesięć?",
			Answers:       []string{"1010", "1100", "1111", "1001"},
			CorrectAnswer: 0,
		},
		{
			Question:      "Który język programowania jest często używany do sztucznej inteligencji?",
			Answers:       []string{"Java", "Python", "C++", "Ruby"},
			CorrectAnswer: 1,
		},
	}

	newGame := &Game{
		Room:     r,
		State:    &GameState{Round: 1},
		Content:  content,
		Category: "",
		IsGame:   true,
		Players:  r.clients,
		mutex:    sync.Mutex{},
	}

	r.game = newGame

	log.Println("new game: ", newGame.Content[0].Question)
	return newGame
}

func (g *Game) NewGameState() *GameState {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	score := g.NewScore()
	log.Println("g.Content[g.State.Round-1].Question", g.Content[g.State.Round-1].Question)
	return &GameState{
		Round:    g.State.Round,
		Question: g.Content[g.State.Round-1].Question,
		Answers:  g.Content[g.State.Round-1].Answers,
		Actions:  []RoundAction{},
		Score:    score,
	}
}

func (g *Game) NewScore() []PlayerScore {
	var score []PlayerScore

	for p := range g.Players {
		playerScore := PlayerScore{
			Name:      p.name,
			Points:    p.points,
			RoundsWon: p.roundsWon,
		}
		score = append(score, playerScore)
	}

	return score
}

func (g *Game) CheckIfShouldBeNextRound() {
	playersInGame := len(g.Players)
	playersFinished := len(g.State.PlayersFinished)
	if playersFinished == playersInGame && playersInGame > 0 {
		g.State.Round++
		newState := g.NewGameState()
		g.State = newState
	}
}

func (g *Game) SendGameState() error {
	log.Println(g.Category, "category send game")
	g.mutex.Lock()
	defer g.mutex.Unlock()

	state := g.State
	stateBytes, err := json.Marshal(state)
	log.Println("Send Game State to the client: ", g.State)
	if err != nil {
		log.Println("Error marshaling game state:", err)
		return err
	}
	event := Event{
		Type:    "update_gamestate",
		Payload: stateBytes,
	}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		log.Println("Error marshaling game state:", err)
		return err
	}
	for client := range g.Room.clients {
		if client == nil {
			return err
		}
		client.receive <- eventBytes
	}
	return nil
}

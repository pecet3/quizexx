package ws

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/pecet3/quizex/external"
)

type Game struct {
	Room       *room
	State      *GameState
	IsGame     bool
	Players    map[*client]string
	Category   string
	Difficulty string
	MaxRounds  int
	Content    []RoundQuestion
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

type RoundQuestion struct {
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
	CorrectAnswer int      `json:"correctAnswer"`
}

func (r *room) GetContentFromGPT() *[]RoundQuestion {
	response, err := external.FetchBodyFromGPT(r.settings.GameCategory, r.settings.Difficulty, r.settings.MaxRounds)
	if err != nil {
		log.Println(err)
	}
	var data *[]RoundQuestion

	err = json.Unmarshal([]byte(response), &data)
	if err != nil {
		log.Println("error with unmarshal data")
	}
	maxRounds, err := strconv.Atoi(r.settings.MaxRounds)

	if err != nil {
		maxRounds = 5
	}
	if len(*data) != maxRounds {
		log.Println("ChatGPT returned insufficient content. Trying to process again...")
		return r.GetContentFromGPT()
	}
	return data
}

func (r *room) CreateGame() *Game {
	log.Println("creating a game")

	data := r.GetContentFromGPT()
	maxRounds, err := strconv.Atoi(r.settings.MaxRounds)

	if err != nil {
		maxRounds = 5
	}
	newGame := &Game{
		Room:       r,
		State:      &GameState{Round: 1},
		IsGame:     false,
		Players:    r.clients,
		Category:   r.settings.GameCategory,
		Difficulty: r.settings.Difficulty,
		MaxRounds:  maxRounds,
		Content:    *data,
	}

	log.Println("created new game in room: ", r.name)
	return newGame
}

func (g *Game) NewGameState() *GameState {

	score := g.NewScore()
	log.Println("created new score in room: ", g.Room.name)

	return &GameState{
		Round:           g.State.Round,
		Question:        g.Content[g.State.Round-1].Question,
		Answers:         g.Content[g.State.Round-1].Answers,
		Actions:         []RoundAction{},
		Score:           score,
		PlayersFinished: []string{},
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

func (g *Game) CheckIfShouldBeNextRound() bool {
	playersInGame := len(g.Players)
	playersFinished := len(g.State.PlayersFinished)
	if playersFinished == playersInGame && playersInGame > 0 {
		return true
	}
	return false
}

func (g *Game) CheckIfIsEndGame() bool {

	isEqualMaxAndCurrentRound := g.State.Round == g.MaxRounds
	isNextRound := g.CheckIfShouldBeNextRound()

	if isEqualMaxAndCurrentRound && !isNextRound {
		log.Println("to jest koniec gry ")
		return true
	}
	log.Println("to nie jest koniec gry ")
	return false
}

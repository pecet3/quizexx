package quiz

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/pecet3/quizex/data/entities"
	"github.com/pecet3/quizex/pkg/logger"
)

type Game struct {
	Room       *Room
	State      *GameState
	IsGame     bool
	Players    map[UUID]*Player
	Category   string
	Difficulty string
	MaxRounds  int
	Content    []RoundQuestion
}

type Player struct {
	user        *entities.User
	isReady     bool
	isSpectator bool
	isAnswered  bool
	answer      int
	points      int
	roundsWon   []uint
	lastActive  time.Time
}

type GameState struct {
	Round           int           `json:"round"`
	Question        string        `json:"question"`
	Answers         []string      `json:"answers"`
	Actions         []RoundAction `json:"actions"`
	Score           []PlayerScore `json:"score"`
	PlayersAnswered []string      `json:"players_answered"`
	RoundWinners    []string      `json:"-"`
}

type RoundAction struct {
	UUID   string `json:"uuid"`
	Answer int    `json:"answer"`
	Round  int    `json:"round"`
}

type PlayerScore struct {
	User       *entities.User `json:"user"`
	Points     int            `json:"points"`
	RoundsWon  []uint         `json:"rounds_won"`
	IsAnswered bool           `json:"is_answered"`
}
type RoundQuestion struct {
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
	CorrectAnswer int      `json:"correct_answer"`
}

func (r *Room) CreateGame() (*Game, error) {
	logger.Info("Creating a game in room: ", r.settings.Name)
	maxRoundsInt := r.settings.MaxRounds
	maxRoundsStr := strconv.Itoa(r.settings.MaxRounds)

	difficulty := r.settings.Difficulty
	category := r.settings.GenContent
	lang := r.settings.Language
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	content, err := fetchQuestionSet(ctx, category, maxRoundsStr, difficulty, lang)
	if err != nil {
		return nil, err
	}
	var questions []RoundQuestion

	err = json.Unmarshal([]byte(content), &questions)
	if err != nil {
		return nil, err
	}

	newGame := &Game{
		Room:       r,
		State:      &GameState{Round: 1},
		IsGame:     false,
		Players:    make(map[UUID]*Player),
		Category:   r.settings.GenContent,
		Difficulty: r.settings.Difficulty,
		MaxRounds:  maxRoundsInt,
		Content:    questions,
	}
	r.game = newGame
	return newGame, nil
}

func (g *Game) newScore() []PlayerScore {
	var score []PlayerScore

	for _, p := range g.Players {
		playerScore := PlayerScore{
			User:      p.user,
			Points:    p.points,
			RoundsWon: p.roundsWon,
		}
		score = append(score, playerScore)
	}

	return score
}

func (g *Game) newGameState(content []RoundQuestion) *GameState {
	logger.Debug("> New Game state in room: ", g.Room.Name)
	g.Content = content
	score := g.newScore()
	return &GameState{
		Round:           g.State.Round,
		Question:        g.Content[g.State.Round-1].Question,
		Answers:         g.Content[g.State.Round-1].Answers,
		Actions:         []RoundAction{},
		Score:           score,
		PlayersAnswered: []string{},
	}
}
func (g *Game) checkIfShouldBeNextRound() bool {
	playersInGame := len(g.Room.clients)
	playersFinished := len(g.State.PlayersAnswered)
	if playersFinished == playersInGame && playersInGame > 0 {
		return true
	}
	return false
}

func (g *Game) checkIfIsEndGame() bool {
	isEqualMaxAndCurrentRound := g.State.Round == g.MaxRounds
	isNextRound := g.checkIfShouldBeNextRound()

	if isEqualMaxAndCurrentRound && isNextRound {
		return true
	}
	return false
}

func (g *Game) checkAnswer(p *Player, action *RoundAction) bool {
	isGoodAnswer := false
	if p.user.UUID == action.UUID {
		p.answer = action.Answer
		if action.Answer == g.Content[g.State.Round-1].CorrectAnswer && !p.isAnswered {
			isGoodAnswer = true
		}
	}
	return isGoodAnswer
}

func (g *Game) toggleClientIsAnswered(p *Player, action *RoundAction) {
	if action.Answer >= 0 && !p.isAnswered {
		g.State.PlayersAnswered = append(g.State.PlayersAnswered, p.user.Name)
		p.isAnswered = true
	}
}
func (g *Game) findWinner() []string {
	highestScore := 0
	for _, c := range g.Room.game.Players {
		if highestScore < c.points {
			highestScore = c.points
			continue
		}
	}
	winners := []string{}
	for _, c := range g.Room.game.Players {
		if highestScore == c.points {
			winners = append(winners, c.user.Name)
			continue
		}
	}
	return winners
}

package ws

type Player struct {
	Name   string
	Answer int
	Points int
}

type GameState struct {
	IsGame          bool
	Category        string
	Round           int
	Question        string
	Answers         []string
	Players         []Player
	PrevRoundWinner []string
}

type QandA struct {
	question      string
	answers       []string
	correctAnswer int
}

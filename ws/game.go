package ws

type Player struct {
	Name   string `json:"name"`
	Answer int    `json:"answer"`
	Points int    `json:"points"`
	Round  int    `json:"round"`
}

type GameState struct {
	IsGame          bool     `json:"isGame"`
	Category        string   `json:"category"`
	Round           int      `json:"round"`
	Question        string   `json:"question"`
	Answers         []string `json:"answers"`
	Players         []Player `json:"players"`
	PrevRoundWinner []string `json:"prevRoundWinner"`
}

type QandA struct {
	question      string
	answers       []string
	correctAnswer int
}

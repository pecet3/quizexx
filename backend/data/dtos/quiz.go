package dtos

type Settings struct {
	Name       string `json:"name"`
	GenContent string `json:"gen_content"`
	Difficulty string `json:"difficulty"`
	MaxRounds  string `json:"max_rounds"`
	Language   string `json:"language"`
}

type Rooms struct {
	Rooms []*Room `json:"rooms"`
}
type Room struct {
	Name       string `json:"name"`
	Players    int    `json:"players"`
	MaxPlayers int    `json:"max_players"`
	Round      int    `json:"round"`
	MaxRounds  int    `json:"max_rounds"`
}

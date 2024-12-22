package dtos

type Settings struct {
	Name       string `json:"name"`
	GenContent string `json:"gen_content"`
	Difficulty string `json:"difficulty"`
	MaxRounds  string `json:"max_rounds"`
	Language   string `json:"language"`
}

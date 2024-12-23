package dtos

import "github.com/go-playground/validator/v10"

type Settings struct {
	Name       string `json:"name" validate:"required"`
	GenContent string `json:"gen_content" validate:"required"`
	Difficulty string `json:"difficulty" validate:"required"`
	MaxRounds  string `json:"max_rounds" validate:"required"`
	Language   string `json:"language" validate:"required"`
}

func (r Settings) Validate(v *validator.Validate) error {
	err := v.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

type Rooms struct {
	Rooms []*Room `json:"rooms"`
}
type Room struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Players    int    `json:"players"`
	MaxPlayers int    `json:"max_players"`
	Round      int    `json:"round"`
	MaxRounds  int    `json:"max_rounds"`
}

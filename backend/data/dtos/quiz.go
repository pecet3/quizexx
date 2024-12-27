package dtos

import "github.com/go-playground/validator/v10"

type Settings struct {
	Name       string `json:"name" validate:"required,alphanumunicode,min=3,max=32"`
	GenContent string `json:"gen_content" validate:"required,alphanumunicode,min=3,max=32"`
	Difficulty string `json:"difficulty" validate:"required,alphanumunicode,min=3,max=32"`
	MaxRounds  string `json:"max_rounds" validate:"required,alphanumunicode,min=1,max=1"`
	Language   string `json:"language" validate:"required,alphanumunicode,min=3,max=32"`
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

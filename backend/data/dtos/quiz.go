package dtos

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Settings struct {
	Name       string `json:"name" validate:"required,alphanumunicode,min=3,max=32"`
	GenContent string `json:"gen_content" validate:"required,min=3,max=64"`
	Difficulty string `json:"difficulty" validate:"required,alphanumunicode,min=3,max=12"`
	MaxRounds  int    `json:"max_rounds" validate:"required,gt=1,lt=10"`
	Language   string `json:"language" validate:"required,alphanumunicode,min=3,max=32"`
}

func (r Settings) Validate(v *validator.Validate) error {
	err := v.Struct(r)
	if !strings.Contains(r.Difficulty, "easy") && !strings.Contains(r.Difficulty, "medium") && !strings.Contains(r.Difficulty, "hard") && !strings.Contains(r.Difficulty, "veryhard") {
		return errors.New("wrong difficulty")
	}
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

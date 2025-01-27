package dtos

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

type Settings struct {
	Name         string `json:"name" validate:"required,min=3,max=32"`
	GenContent   string `json:"gen_content" validate:"required,min=3,max=64"`
	Difficulty   string `json:"difficulty" validate:"required,oneof=easy medium hard veryhard"`
	MaxRounds    int    `json:"max_rounds" validate:"required,gt=1,lt=10"`
	Language     string `json:"language" validate:"required,alphanumunicode,min=3,max=32"`
	SecForAnswer int    `json:"sec_for_answer" validate:"required,oneof=10 15 20 30 45 60"`
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

func (r Room) Send(w io.Writer) error {
	if err := json.NewEncoder(w).Encode(&r); err != nil {
		return err
	}
	return nil
}

type FunFact struct {
	Topic   string `json:"topic"`
	Content string `json:"content"`
}

func (f FunFact) Send(w io.Writer) error {
	if err := json.NewEncoder(w).Encode(&f); err != nil {
		return err
	}
	return nil
}

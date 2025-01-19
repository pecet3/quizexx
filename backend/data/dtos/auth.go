package dtos

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

type Register struct {
	Name  string `json:"name" validate:"required,min=4,alphanumunicode"`
	Email string `json:"email" validate:"required,email"`
}

func (r Register) Validate(v *validator.Validate) error {
	err := v.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

type Exchange struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required,min=6,max=6"`
}

func (r Exchange) Validate(v *validator.Validate) error {
	err := v.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

type Login struct {
	Email string `json:"email" validate:"required,email"`
}

func (r *Login) Validate(v *validator.Validate) error {
	err := v.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

func (r Login) Send(w io.Writer) error {
	if err := json.NewEncoder(w).Encode(&r); err != nil {
		return err
	}
	return nil
}

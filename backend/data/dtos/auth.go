package dtos

import (
	"github.com/go-playground/validator/v10"
)

type RegisterDTO struct {
	Name  string `json:"name" validate:"required,min=4"`
	Email string `json:"email" validate:"required,email"`
}

func (r RegisterDTO) Validate(v *validator.Validate) error {
	err := v.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

type ExchangeDTO struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required,min=6,max=6"`
}

func (r ExchangeDTO) Validate(v *validator.Validate) error {
	err := v.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

type LoginDTO struct {
	Email string `json:"email" validate:"required,email"`
}

func (r LoginDTO) Validate(v *validator.Validate) error {
	err := v.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

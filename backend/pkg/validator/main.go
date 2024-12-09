package validator

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	V *validator.Validate
}
type Validatable interface {
	Validate(v *validator.Validate) error
}

func New() *Validator {
	return &Validator{
		V: validator.New(),
	}
}

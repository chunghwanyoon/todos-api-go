package utils

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validator *validator.Validate
}

func (c *Validator) Validate(i interface{}) error {
	return c.Validator.Struct(i)
}

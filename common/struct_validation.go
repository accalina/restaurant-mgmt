package common

import "github.com/go-playground/validator/v10"

func Validate(f interface{}) error {
	validate := validator.New()
	return validate.Struct(f)
}

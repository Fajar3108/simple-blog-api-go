package validation

import "github.com/go-playground/validator"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validate[T any](request T) error {
	return validate.Struct(request)
}

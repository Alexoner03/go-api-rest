package utils

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func CheckErrorAndDie(err error) {
	if err != nil {
		panic(err)
	}
}

func ValidateStruct(model interface{}) []*ErrorResponse {
	var validate = validator.New()
	var errors []*ErrorResponse
	err := validate.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

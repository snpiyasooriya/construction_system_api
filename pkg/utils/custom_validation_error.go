package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

func CustomValidationErrors(validateStruct interface{}) []ValidationError {
	// Validate the input
	validate := validator.New()
	if err := validate.Struct(validateStruct); err != nil {

		var errorMessages []ValidationError
		for _, err := range err.(validator.ValidationErrors) {

			switch err.Tag() {
			case "required":
				errorMessages = append(errorMessages, ValidationError{
					err.Field(),
					fmt.Sprintf("%s is required", err.Field()),
				})
			case "max":
				errorMessages = append(errorMessages, ValidationError{
					err.Field(),
					fmt.Sprintf("%s cannot be longer than %s characters", err.Field(), err.Param()),
				})
			case "min":
				errorMessages = append(errorMessages, ValidationError{
					err.Field(),
					fmt.Sprintf("%s  must be at least %s characters long", err.Field(), err.Param()),
				})
			case "email":
				errorMessages = append(errorMessages, ValidationError{
					err.Field(),
					fmt.Sprintf("%s must be a valid email address", err.Field()),
				})

			}
		}
		return errorMessages
	}
	return nil
}

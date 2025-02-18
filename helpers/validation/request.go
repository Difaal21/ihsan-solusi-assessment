package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type InvalidSchema struct {
	Field   string `json:"field"`
	Message any    `json:"message"`
}

func RequestBody(handler *validator.Validate, body interface{}) interface{} {
	err := handler.Struct(body)
	if err == nil {
		return nil
	}

	// Check if the error is of type validator.ValidationErrors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var invalid []*InvalidSchema

		for _, errorField := range validationErrors {
			invalid = append(invalid, &InvalidSchema{
				Field:   errorField.Field(),
				Message: fmt.Sprintf("invalid '%s' with value '%v'", errorField.Field(), errorField.Value()),
			})
		}
		return invalid
	}

	if invalidValidationError, ok := err.(*validator.InvalidValidationError); ok {
		return []*InvalidSchema{
			{
				Field:   "InvalidValidationError",
				Message: invalidValidationError.Error(),
			},
		}
	}

	return []*InvalidSchema{
		{
			Field:   "UnknownError",
			Message: err.Error(),
		},
	}
}

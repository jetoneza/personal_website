package models

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	Error bool        `json:"error"`
	Field string      `json:"field"`
	Tag   string      `json:"tag"`
	Value interface{} `json:"value,omitempty"`
}

var validate = validator.New()

func Validate(data interface{}) []*ErrorResponse {
	var validationErrors []*ErrorResponse

	errors := validate.Struct(data)

	if errors != nil {
		for _, err := range errors.(validator.ValidationErrors) {
			var e ErrorResponse

			e.Field = err.Field()
			e.Tag = err.Tag()
			e.Value = err.Value()
			e.Error = true

			validationErrors = append(validationErrors, &e)
		}
	}

	return validationErrors
}

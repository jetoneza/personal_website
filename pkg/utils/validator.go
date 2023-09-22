package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Error bool        `json:"error"`
	Field string      `json:"field"`
	Tag   string      `json:"tag"`
	Value interface{} `json:"value,omitempty"`
}

var validate = validator.New()

func Validate(data interface{}) *fiber.Error {
	errors := validate.Struct(data)

	if errors != nil {
		var err []string

		for _, e := range errors.(validator.ValidationErrors) {
			err = append(
				err,
				fmt.Sprintf("`%v` with value `%v` doesn't satisfy the `%v` constraint", e.Field(), e.Value(), e.Tag()),
			)
		}

		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: strings.Join(err, ","),
		}
	}

	return nil
}

// -----------------------
// Custom validation rules
// -----------------------

// Password validation rule: required, min=6, max=100
var _ = validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
	l := len(fl.Field().String())

	return l >= 6 && l <= 100
})

package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/guemidiborhane/explore-go/errors"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	Validator struct {
		validator *validator.Validate
	}
)

var validate = validator.New(validator.WithRequiredStructEnabled())

var Validation *Validator

func (v Validator) Validate(data interface{}) []ErrorResponse {
	errors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			errors = append(errors, elem)
		}
	}

	return errors
}

func Validate(c *fiber.Ctx, body interface{}) error {
	c.BodyParser(&body)

	if err := Validation.Validate(body); len(err) > 0 && err[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range err {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%s'",
				err.FailedField,
				err.Tag,
			))
		}

		return errors.BadRequest(strings.Join(errMsgs, ", "))
	}

	return c.Next()
}

func Setup() {
	Validation = &Validator{
		validator: validate,
	}
}

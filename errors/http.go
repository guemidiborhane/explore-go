package errors

import (
	"github.com/gofiber/fiber/v2"
)

func HandleHttpErrors(ctx *fiber.Ctx, err error) error {
	if e, ok := err.(*HttpError); ok {
		return ctx.Status(e.Status).JSON(e)
	} else if e, ok := err.(*fiber.Error); ok {
		return ctx.Status(e.Code).JSON(HttpError{Status: e.Code, Code: "internal-server", Message: e.Message})
	} else {
		return ctx.Status(499).JSON(HttpError{Status: 500, Code: "internal-server", Message: err.Error()})
	}
}

type HttpError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *HttpError) Error() string {
	return e.Message
}

func EntityNotFound(m string) *HttpError {
	return &HttpError{Status: 404, Code: "entity-not-found", Message: m}
}

func BadRequest(m string) *HttpError {
	return &HttpError{Status: 400, Code: "bad-request", Message: m}
}

func Unexpected(m string) *HttpError {
	return &HttpError{Status: 500, Code: "internal-server", Message: m}
}

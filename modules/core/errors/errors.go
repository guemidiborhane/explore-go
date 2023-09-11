package errors

import (
	"github.com/gofiber/fiber/v2"
)

func Handle(ctx *fiber.Ctx, err error) error {
	if e, ok := err.(*Error); ok {
		return ctx.Status(e.Status).JSON(e)
	} else if e, ok := err.(*fiber.Error); ok {
		return ctx.Status(e.Code).JSON(Error{Status: e.Code, Code: "internal-server", Message: e.Message})
	} else {
		return ctx.Status(499).JSON(Error{Status: 500, Code: "internal-server", Message: err.Error()})
	}
}

type Error struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func EntityNotFound(m string) *Error {
	return &Error{Status: 404, Code: "entity-not-found", Message: m}
}

func BadRequest(m string) *Error {
	return &Error{Status: 400, Code: "bad-request", Message: m}
}

func Unexpected(m string) *Error {
	return &Error{Status: 500, Code: "internal-server", Message: m}
}

package auth

import (
	"github.com/gofiber/fiber/v2"
	"explore-go/config/validator"
	"explore-go/errors"
)

func validateUser(c *fiber.Ctx) error {
	return validator.Validate(c, new(User))
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func validateLogin(c *fiber.Ctx) error {
	return validator.Validate(c, new(LoginRequest))
}

func Auth(c *fiber.Ctx) error {
	session, err := store.Get(c)

	if err != nil {
		return errors.Unauthorized
	}

	if session.Get(AUTH_KEY) == nil {
		return errors.Unauthorized
	}

	return c.Next()
}

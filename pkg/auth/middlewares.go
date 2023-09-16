package auth

import (
	"explore-go/config/validator"
	"explore-go/errors"

	"github.com/gofiber/fiber/v2"
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

func CheckAuthenticated(c *fiber.Ctx) error {
	session, err := store.Get(c)
	if err != nil {
		return errors.Unauthorized
	}

	if session.Get(AUTH_KEY) == nil {
		return errors.Unauthorized
	}

	c.Locals(USER_ID, session.Get(USER_ID).(uint))

	return c.Next()
}

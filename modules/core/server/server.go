package server

import (
	"core/errors"
	"core/middlewares"

	"github.com/gofiber/fiber/v2"
)

var Application *fiber.App

func Setup() {
	Application = fiber.New(fiber.Config{
		Prefork:      true,
		ErrorHandler: errors.Handle,
	})

	middlewares.Setup(Application)
}

package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guemidiborhane/explore-go/errors"
)

var App *fiber.App

func Setup() {
	App = fiber.New(fiber.Config{
		Prefork:      true,
		ErrorHandler: errors.HandleHttpErrors,
	})
}

func Start() {
	SetupMiddlewares()
	httpListen()
}

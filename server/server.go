package server

import (
	"explore-go/errors"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

var App *fiber.App

func Setup() {
	App = fiber.New(fiber.Config{
		Prefork:      true,
		ErrorHandler: errors.HandleHttpErrors,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})
}

func Start() {
	SetupMiddlewares(App)
	httpListen()
}

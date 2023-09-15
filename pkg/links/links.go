package links

import (
	"github.com/gofiber/fiber/v2"
	"explore-go/pkg/setup"
)

var app *fiber.App

func Setup(args *setup.SetupArgs) {
	database = args.Database
	app = args.Application
	router = *args.Router

	setupModels()
	setupRoutes()
}

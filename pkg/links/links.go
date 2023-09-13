package links

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guemidiborhane/explore-go/pkg/setup"
)

var app *fiber.App

func setupModels() {
	database.AutoMigrate(&Link{})
}

func Setup(args *setup.SetupArgs) {
	database = args.Database
	app = args.Application
	router = *args.Router

	setupModels()
	setupRoutes()
}

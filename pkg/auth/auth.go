package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"explore-go/pkg/setup"
)

var (
	app   *fiber.App
	store *session.Store
)

func Setup(args *setup.SetupArgs) {
	database = args.Database
	app = args.Application
	router = *args.Router
	store = args.Session

	setupRoutes()
	setupModels()
}

package router

import (
	"core/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

var ApiRouter fiber.Router

func Setup() {
	ApiRouter = server.Application.Group("/api", logger.New())
	ApiRouter.Get("/monitor", monitor.New())
}

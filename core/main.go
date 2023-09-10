package core

import (
	application "core/config"
	"core/config/initializers"

	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Setup() {
	initializers.InitDB()
	initializers.InitServer()

	application.Fiber.Get("/monitor", monitor.New())
}

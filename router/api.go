package router

import (
	"explore-go/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var ApiRouter fiber.Router

func SetupAPI() {
	ApiRouter = server.App.Group("/api", logger.New())
	ApiRouter.Use(func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		return c.Next()
	})

	ApiRouter.Use(
		HelmetMiddleware,
		CorsMiddleware,
		CsrfMiddleware,
		RecoverMiddleware,
		CompressMiddleware,
		EtagMiddleware,
	)
}

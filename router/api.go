package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/guemidiborhane/explore-go/server"
)

var ApiRouter fiber.Router

func Setup() {
	ApiRouter = server.App.Group("/api", logger.New())
	ApiRouter.Use(func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		return c.Next()
	})
	server.SetupMiddlewares(ApiRouter)
}

package routes

import (
	"core/server"

	"github.com/gofiber/fiber/v2"
)

func Setup() {
	server.Application.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
}

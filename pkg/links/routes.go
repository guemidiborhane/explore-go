package links

import (
	"github.com/gofiber/fiber/v2"
)

var router fiber.Router

func setupRoutes() {
	group := router.Group("/links")

	group.Get("/", Index)
	group.Get("/new", Build)
	group.Post("/", New)
	group.Get("/:id", Show)
	group.Patch("/:id", Edit)
	group.Delete("/:id", Delete)

	app.Get("/r/:short", Redirect)
}

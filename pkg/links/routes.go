package links

import (
	"explore-go/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

var router fiber.Router

func setupRoutes() {
	group := router.Group("/links", auth.CheckAuthenticated)

	group.Get("/", Index)
	group.Get("/new", Build)
	group.Post("/", validateLink, New)
	group.Get("/:id", Show)
	group.Patch("/:id", validateLink, Edit)
	group.Delete("/:id", Delete)

	app.Get("/r/:short", Redirect)
}

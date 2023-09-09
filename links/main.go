package links

import (
	application "core/config"
	"links/handlers"
	"links/models"
)

func setupModels() {
	application.Database.AutoMigrate(&models.Link{})
}

func setupRoutes() {
	group := application.Router.Group("/links")

	group.Get("/", handlers.Index)
	group.Post("/", handlers.Create)
	group.Get("/:id", handlers.Show)
	group.Patch("/:id", handlers.Update)
	group.Delete("/:id", handlers.Destroy)

	application.Fiber.Get("/r/:short", handlers.Redirect)
}

func Setup() {
	setupModels()
	setupRoutes()
}

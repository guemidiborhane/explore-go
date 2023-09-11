package links

import (
	"core/database"
	"core/router"
	"core/server"
	"links/handlers"
	"links/models"
)

func setupModels() {
	database.Database.AutoMigrate(&models.Link{})
}

func setupRoutes() {
	group := router.ApiRouter.Group("/links")

	group.Get("/", handlers.Index)
	group.Get("/new", handlers.New)
	group.Post("/", handlers.Create)
	group.Get("/:id", handlers.Show)
	group.Patch("/:id", handlers.Update)
	group.Delete("/:id", handlers.Destroy)

	server.Application.Get("/r/:short", handlers.Redirect)
}

func Setup() {
	setupModels()
	setupRoutes()
}

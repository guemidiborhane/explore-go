package links

import (
	application "github.com/guemidiborhane/explore-go/config"
	handlers "github.com/guemidiborhane/explore-go/links/handlers"
	models "github.com/guemidiborhane/explore-go/links/models"
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
}

func Setup() {
	setupModels()
	setupRoutes()
}

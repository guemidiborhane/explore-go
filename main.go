package main

import (
	application "github.com/guemidiborhane/explore-go/config"
	"github.com/guemidiborhane/explore-go/config/initializers"
	linksHandler "github.com/guemidiborhane/explore-go/links/handlers"
)

func main() {
	initializers.InitDB()
	initializers.InitServer()
	linksHandler.SetupRoutes()

	application.Fiber.Listen(":3000")
}

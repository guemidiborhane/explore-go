package main

import (
	application "github.com/guemidiborhane/explore-go/config"
	"github.com/guemidiborhane/explore-go/config/initializers"
	"github.com/guemidiborhane/explore-go/links"
)

func main() {
	initializers.InitDB()
	initializers.InitServer()
	links.Setup()

	application.Fiber.Listen(":3000")
}

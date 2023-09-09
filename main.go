package main

import (
	"core"
	"links"

	application "core/config"
)

func main() {
	core.Setup()
	links.Setup()

	application.Fiber.Listen(":3000")
}

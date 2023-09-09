package main

import (
	"core"
	"fmt"
	"links"

	application "core/config"
	"core/utils"

	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("Error loading .env file")
	}
	core.Setup()
	links.Setup()

	application.Fiber.Listen(
		fmt.Sprintf(
			"%s:%d",
			utils.GetEnv("HOST", "0.0.0.0"),
			utils.ParseUint(utils.GetEnv("PORT", "3000"), 64),
		),
	)
}

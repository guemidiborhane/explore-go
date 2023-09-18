package main

import (
	"explore-go/config"
	"explore-go/config/validator"
	"explore-go/database"
	"explore-go/database/session"
	"explore-go/pkg"
	"explore-go/router"
	"explore-go/server"
	"explore-go/static"
	"explore-go/websocket"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Load()
	server.Setup()
	database.Setup()
	session.Setup()
	validator.Setup()
	websocket.Setup()

	// Required to run first since it registers the group
	// all other packages are gonna register their routes on
	router.Setup()
	pkg.Setup()

	if !fiber.IsChild() {
		go func() {
			// wait 1 second before starting
			time.Sleep(3 * time.Second)
			websocket.Send(websocket.Message{
				Channel: "system",
				Message: "reload",
			})
		}()
	}

	static.Setup()
	server.Start()
}

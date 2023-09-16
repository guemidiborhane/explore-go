package main

import (
	"explore-go/config"
	"explore-go/config/validator"
	"explore-go/database"
	"explore-go/database/session"
	"explore-go/pkg"
	"explore-go/router"
	"explore-go/server"
	"explore-go/server/websocket"

	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	config.Load()
	server.Setup()
	database.Setup()
	session.Setup()
	validator.Setup()

	// Required to run first since it registers the group
	// all other packages are gonna register their routes on
	router.Setup()
	pkg.Setup()
	router.ApiRouter.Use("/monitor", monitor.New())
	websocket.SetupWebsocketServer()

	websocket.Send(websocket.Message{
		Channel: "system",
		Message: "reload",
	})

	SetupStatic()
	server.Start()
}

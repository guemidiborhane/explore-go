package router

import "github.com/gofiber/fiber/v2/middleware/monitor"

func Setup() {
	SetupAPI()
	SetupWebsocketRoutes()

	ApiRouter.Use("/monitor", monitor.New())
}

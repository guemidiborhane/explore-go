package router

import (
	ws "explore-go/websocket"

	"github.com/gofiber/contrib/websocket"
)

func SetupWebsocketRoutes() {
	ws.Websocket.Router.Use(HelmetMiddleware, CompressMiddleware, RecoverMiddleware, EtagMiddleware)
	ws.Websocket.Router.Get("/", ws.UpgradeHandler)
	ws.Websocket.Router.Get("/", websocket.New(ws.WsHandler), websocket.New(ws.DisconnectHandler))
}

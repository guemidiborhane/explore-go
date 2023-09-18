package websocket

import (
	"context"
	"explore-go/pkg/auth"
	"explore-go/server"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Message struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
}

var (
	ClientsMutex sync.Mutex
	Clients      = make(map[*websocket.Conn]struct{})
	mutex        sync.Mutex
	ctx          = context.Background()
)

type WS struct {
	Router fiber.Router
}

var Websocket *WS

func Setup() {
	Websocket = &WS{
		Router: server.App.Group("/ws", auth.CheckAuthenticated, logger.New()),
	}
}

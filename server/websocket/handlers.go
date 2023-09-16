package websocket

import (
	"encoding/json"
	"explore-go/database"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func upgradeHandler(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)

		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func wsHandler(c *websocket.Conn) {
	mutex.Lock()
	Clients[c] = struct{}{}
	mutex.Unlock()

	pubsub := database.Storage.Conn().Subscribe(ctx, "websocket")

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			log.Println("pubsub error:", err)
			break
		}
		var message Message
		if err := json.Unmarshal([]byte(msg.Payload), &message); err != nil {
			log.Println("json unmarshal error:", err)
			continue
		}
		if err := c.WriteJSON(message); err != nil {
			log.Println("write:", err)
		}
	}
}

package websocket

import (
	"encoding/json"
	"explore-go/database"
	"log"
)

func Send(msg Message) {
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println("json marshal error:", err)
		return
	}

	if err := database.Storage.Conn().Publish(ctx, "websocket", payload).Err(); err != nil {
		log.Println("publish error:", err)
	} else {
		log.Println("Message published to Redis:", string(payload))
	}
}

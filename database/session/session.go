package session

import (
	"explore-go/database"
	"explore-go/utils"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var Session *session.Store

func Setup() {
	Session = session.New(session.Config{
		Storage:        database.Storage,
		CookieHTTPOnly: true,
		Expiration:     24 * time.Hour,
		KeyLookup:      "cookie:session_id",
		KeyGenerator:   utils.RandomID,
	})
}

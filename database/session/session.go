package session

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/guemidiborhane/explore-go/database"
	"github.com/guemidiborhane/explore-go/utils"
)

var Session *session.Store

func Setup() {
	Session = session.New(session.Config{
		Storage:        database.Storage,
		CookieHTTPOnly: true,
		Expiration:     24 * time.Hour,
		KeyLookup:      "cookie:session_id",
		KeyGenerator:   utils.UUIDv4,
	})
}

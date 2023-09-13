package session

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"github.com/guemidiborhane/explore-go/utils"
)

var redisConfig = &redis.Config{
	Host:     utils.GetEnv("REDIS_HOST", "localhost"),
	Port:     utils.ParseInt(utils.GetEnv("REDIS_PORT", "6379")),
	Username: utils.GetEnv("REDIS_USERNAME", ""),
	Password: utils.GetEnv("REDIS_PASSWORD", ""),
	Database: utils.ParseInt(utils.GetEnv("REDIS_DATABASE", "0")),
}

var (
	Session *session.Store
	storage *redis.Storage = redis.New(*redisConfig)
)

func Setup() {
	Session = session.New(session.Config{
		Storage:        storage,
		CookieHTTPOnly: true,
		Expiration:     24 * time.Hour,
		KeyLookup:      "cookie:session_id",
		KeyGenerator:   utils.UUIDv4,
	})
}

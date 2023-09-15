package database

import (
	"github.com/gofiber/storage/redis"
	"github.com/guemidiborhane/explore-go/utils"
)

var redisConfig = redis.Config{
	Host:     utils.GetEnv("REDIS_HOST", "localhost"),
	Port:     utils.ParseInt(utils.GetEnv("REDIS_PORT", "6379")),
	Username: utils.GetEnv("REDIS_USERNAME", ""),
	Password: utils.GetEnv("REDIS_PASSWORD", ""),
	Database: utils.ParseInt(utils.GetEnv("REDIS_DATABASE", "0")),
}

var Storage *redis.Storage = redis.New(redisConfig)

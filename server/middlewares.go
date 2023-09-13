package server

import (
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupMiddlewares() {
	App.Use(helmet.New())
	App.Use(recover.New())
	App.Use(cors.New())
	App.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	App.Use(etag.New())
}

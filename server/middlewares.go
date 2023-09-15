package server

import (
	"explore-go/database"
	"explore-go/errors"
	"explore-go/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupMiddlewares(router fiber.Router) {
	router.Use(helmet.New())
	router.Use(cors.New())
	router.Use(csrf.New(csrf.Config{
		Storage:        database.Storage,
		CookieSameSite: "Strict",
		KeyGenerator:   utils.RandomID,
		ErrorHandler:   errors.HandleHttpErrors,
	}))
	router.Use(recover.New())
	router.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	router.Use(etag.New())
}

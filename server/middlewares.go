package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/utils"
	"github.com/guemidiborhane/explore-go/database"
	"github.com/guemidiborhane/explore-go/errors"
)

func SetupMiddlewares(router fiber.Router) {
	router.Use(helmet.New())
	router.Use(csrf.New(csrf.Config{
		Storage:        database.Storage,
		KeyLookup:      "cookie:csrf_",
		CookieName:     "csrf_",
		CookieSameSite: "Strict",
		CookieHTTPOnly: true,
		Expiration:     5 * time.Minute,
		KeyGenerator:   utils.UUID,
		ErrorHandler:   errors.HandleHttpErrors,
	}))
	router.Use(recover.New())
	router.Use(cors.New())
	router.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	router.Use(etag.New())
}

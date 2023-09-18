package router

import (
	"explore-go/database"
	"explore-go/errors"
	"explore-go/utils"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var CsrfMiddleware = csrf.New(csrf.Config{
	Storage:        database.Storage,
	CookieSameSite: "Strict",
	KeyGenerator:   utils.RandomID,
	ErrorHandler:   errors.HandleHttpErrors,
})

var HelmetMiddleware = helmet.New()
var CorsMiddleware = cors.New(cors.Config{})
var RecoverMiddleware = recover.New()
var CompressMiddleware = compress.New(compress.Config{
	Level: compress.LevelBestSpeed,
})

var EtagMiddleware = etag.New()
var LoggerMiddleware = logger.New()

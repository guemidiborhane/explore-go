package application

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var Fiber *fiber.App
var Router fiber.Router
var Database *gorm.DB

package setup

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type SetupArgs struct {
	Application *fiber.App
	Router      *fiber.Router
	Database    *gorm.DB
}

package setup

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type SetupArgs struct {
	Application  *fiber.App
	Router       *fiber.Router
	Database     *gorm.DB
	SessionStore *session.Store
}

package links

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guemidiborhane/explore-go/config/validator"
)

func validateLink(c *fiber.Ctx) error {
	return validator.Validate(c, new(Link))
}

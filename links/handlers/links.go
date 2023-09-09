package handlers

import (
	"links/models"
	"links/queries"

	"core/utils"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	links, err := queries.All()

	if err != nil {
		return utils.HandleError(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(links)
}

func Show(c *fiber.Ctx) error {
	id := utils.ParseUint(c.Params("id"), 64)
	link, err := queries.Get(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(link)
}

func Create(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var link models.Link
	link.Short = utils.RandomShort(8)

	if err := c.BodyParser(&link); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err,
		})
	}

	if err := queries.Create(&link); err != nil {
		return utils.HandleError(err, c)
	}

	return c.Status(fiber.StatusCreated).JSON(link)
}

func Update(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var link models.Link
	link, err := queries.Get(utils.ParseUint(c.Params("id"), 64))

	if err != nil {
		return utils.HandleError(err, c)
	}

	if err := c.BodyParser(&link); err != nil {
		return utils.HandleError(err, c)
	}

	if err := queries.Update(link); err != nil {
		return utils.HandleError(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(link)
}

func Destroy(c *fiber.Ctx) error {
	id := utils.ParseUint(c.Params("id"), 64)
	link, err := queries.Get(id)

	if err != nil {
		return utils.HandleError(err, c)
	}

	if err := queries.Destroy(link); err != nil {
		return utils.HandleError(err, c)
	}

	return c.Status(fiber.StatusNoContent).JSON(link)
}

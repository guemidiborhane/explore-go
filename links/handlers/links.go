package handlers

import (
	"links/models"
	"links/queries"

	helpers "core/utils"
	"links/utils"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	links, err := queries.All()

	if err != nil {
		return helpers.HandleError(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(links)
}

func Show(c *fiber.Ctx) error {
	var link models.Link
	id := helpers.ParseUint(c.Params("id"), 64)

	if err := queries.Get(&link, id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
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
			"message": err.Error(),
		})
	}

	if err := queries.Create(&link); err != nil {
		return helpers.HandleError(err, c)
	}

	return c.Status(fiber.StatusCreated).JSON(link)
}

func Update(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var link models.Link
	id := helpers.ParseUint(c.Params("id"), 64)
	if err := queries.Get(&link, id); err != nil {
		return helpers.HandleError(err, c)
	}

	if err := c.BodyParser(&link); err != nil {
		return helpers.HandleError(err, c)
	}

	if err := queries.Update(&link); err != nil {
		return helpers.HandleError(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(link)
}

func Destroy(c *fiber.Ctx) error {
	var link models.Link
	id := helpers.ParseUint(c.Params("id"), 64)
	if err := queries.Get(&link, id); err != nil {
		return helpers.HandleError(err, c)
	}

	if err := queries.Destroy(link); err != nil {
		return helpers.HandleError(err, c)
	}

	return c.Status(fiber.StatusNoContent).JSON(link)
}

func Redirect(c *fiber.Ctx) error {
	var link models.Link
	short := c.Params("short")
	if err := queries.GetByShort(&link, short); err != nil {
		return helpers.HandleError(err, c)
	}

	return c.Status(fiber.StatusTemporaryRedirect).Redirect(link.Link)
}

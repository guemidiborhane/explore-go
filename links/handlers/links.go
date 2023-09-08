package handlers

import (
	"github.com/gofiber/fiber/v2"
	application "github.com/guemidiborhane/explore-go/config"
	"github.com/guemidiborhane/explore-go/links/models"
	"github.com/guemidiborhane/explore-go/links/queries"
	"github.com/guemidiborhane/explore-go/utils"
)

func SetupRoutes() {
	group := application.Router.Group("/links")

	group.Get("/", index)
	group.Post("/", create)
	group.Get("/:id", show)
	group.Patch("/:id", update)
	group.Delete("/:id", destroy)
}

func index(c *fiber.Ctx) error {
	links, err := queries.GetAllLinks()

	if err != nil {
		return utils.HandleError(err, c)
	}

	return c.Status(fiber.StatusFound).JSON(links)
}

func show(c *fiber.Ctx) error {
	id := utils.ParseUint(c.Params("id"), 64)
	link, err := queries.GetLink(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusFound).JSON(link)
}

func create(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var link models.Link
	link.Short = utils.RandomShort(8)

	if err := c.BodyParser(&link); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err,
		})
	}

	if err := queries.CreateLink(&link); err != nil {
		return utils.HandleError(err, c)
	}

	return c.Status(fiber.StatusCreated).JSON(link)
}

func update(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var link models.Link
	link, err := queries.GetLink(utils.ParseUint(c.Params("id"), 64))

	if err != nil {
		return utils.HandleError(err, c)
	}

	if err := c.BodyParser(&link); err != nil {
		return utils.HandleError(err, c)
	}

	if err := queries.UpdateLink(link); err != nil {
		return utils.HandleError(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(link)
}

func destroy(c *fiber.Ctx) error {
	id := utils.ParseUint(c.Params("id"), 64)
	link, err := queries.GetLink(id)

	if err != nil {
		return utils.HandleError(err, c)
	}

	if err := queries.DestroyLink(link); err != nil {
		return utils.HandleError(err, c)
	}

	return c.Status(fiber.StatusNoContent).JSON(link)
}

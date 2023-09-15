package links

import (
	"explore-go/errors"
	"explore-go/utils"

	"github.com/gofiber/fiber/v2"
)

var (
	NotFoundError   *errors.HttpError = errors.EntityNotFound("Link not found")
	BadRequestError *errors.HttpError = errors.BadRequest("Invalid parameters")
)

type ResponseBody struct {
	ID    uint   `json:"id"`
	Link  string `json:"link"`
	Short string `json:"short"`
}

func Index(c *fiber.Ctx) error {
	var links []Link
	if err := All(&links); err != nil {
		return err
	}

	responses := []ResponseBody{}

	for _, link := range links {
		responses = append(responses, link.JSONResponse())
	}

	return c.Status(fiber.StatusOK).JSON(responses)
}

func Show(c *fiber.Ctx) error {
	link := Link{
		ID: uint(utils.ParseUint(c.Params("id"), 64)),
	}
	if err := link.Get(); err != nil {
		return NotFoundError
	}

	return c.Status(fiber.StatusOK).JSON(link.JSONResponse())
}

func Build(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(Link{})
}

func New(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var link Link
	link.Short = utils.Random(8)

	if err := c.BodyParser(&link); err != nil {
		return BadRequestError
	}

	if err := link.Create(); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(link.JSONResponse())
}

func Edit(c *fiber.Ctx) error {
	c.Accepts("application/json")

	link := Link{
		ID: uint(utils.ParseUint(c.Params("id"), 64)),
	}
	if err := link.Get(); err != nil {
		return NotFoundError
	}

	if err := c.BodyParser(&link); err != nil {
		return BadRequestError
	}

	if err := link.Update(); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(link.JSONResponse())
}

func Delete(c *fiber.Ctx) error {
	link := Link{
		ID: uint(utils.ParseUint(c.Params("id"), 64)),
	}
	if err := link.Get(); err != nil {
		return NotFoundError
	}

	if err := link.Destroy(); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(link.JSONResponse())
}

func Redirect(c *fiber.Ctx) error {
	link := Link{
		Short: c.Params("short"),
	}

	if err := link.GetByShort(); err != nil {
		return NotFoundError
	}

	return c.Status(fiber.StatusTemporaryRedirect).Redirect(link.Link)
}

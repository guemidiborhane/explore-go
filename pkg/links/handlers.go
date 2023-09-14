package links

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guemidiborhane/explore-go/errors"
	"github.com/guemidiborhane/explore-go/utils"
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
		responses = append(responses, response(&link))
	}

	return c.Status(fiber.StatusOK).JSON(responses)
}

func Show(c *fiber.Ctx) error {
	link := Link{
		ID: uint(utils.ParseUint(c.Params("id"), 64)),
	}
	if err := Get(&link); err != nil {
		return NotFoundError
	}

	return c.Status(fiber.StatusOK).JSON(response(&link))
}

func Build(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(Link{})
}

func New(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var link Link
	link.Short = RandomShort(8)

	if err := c.BodyParser(&link); err != nil {
		return BadRequestError
	}

	if err := Create(&link); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(response(&link))
}

func Edit(c *fiber.Ctx) error {
	c.Accepts("application/json")

	link := Link{
		ID: uint(utils.ParseUint(c.Params("id"), 64)),
	}
	if err := Get(&link); err != nil {
		return NotFoundError
	}

	if err := c.BodyParser(&link); err != nil {
		return BadRequestError
	}

	if err := Update(&link); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(response(&link))
}

func Delete(c *fiber.Ctx) error {
	link := Link{
		ID: uint(utils.ParseUint(c.Params("id"), 64)),
	}
	if err := Get(&link); err != nil {
		return NotFoundError
	}

	if err := Destroy(link); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(response(&link))
}

func Redirect(c *fiber.Ctx) error {
	link := Link{
		Short: c.Params("short"),
	}

	if err := GetByShort(&link); err != nil {
		return NotFoundError
	}

	return c.Status(fiber.StatusTemporaryRedirect).Redirect(link.Link)
}

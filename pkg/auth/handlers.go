package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guemidiborhane/explore-go/errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	NotFoundError   *errors.HttpError = errors.EntityNotFound("User not found")
	BadRequestError *errors.HttpError = errors.BadRequest("Invalid parameters")
	AUTH_KEY        string            = "authenticated"
	USER_ID         string            = "user_id"
)

type ResponseBody struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

func Register(c *fiber.Ctx) error {
	var user User

	if err := c.BodyParser(&user); err != nil {
		return BadRequestError
	}

	if err := Create(&user); err != nil {
		return errors.Unexpected(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(response(&user))
}

func Show(c *fiber.Ctx) error {
	var user User

	if err := GetCurrentUser(c, &user); err != nil {
		return err
	}

	return c.JSON(response(&user))
}

func Login(c *fiber.Ctx) error {
	var body LoginRequest

	if err := c.BodyParser(&body); err != nil {
		return BadRequestError
	}

	user := User{
		Username: body.Username,
	}

	if err := GetByUsername(&user); err != nil {
		return NotFoundError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return errors.Unauthorized
	}

	setSession(c, &user)

	return c.Status(fiber.StatusOK).JSON(response(&user))
}

func Logout(c *fiber.Ctx) error {
	session, err := store.Get(c)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	if err = session.Destroy(); err != nil {
		return errors.Unexpected(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Goodbye!",
	})
}

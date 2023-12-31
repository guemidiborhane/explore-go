package auth

import (
	"explore-go/errors"

	"github.com/gofiber/fiber/v2"
)

func GetCurrentUser(c *fiber.Ctx, user *User) error {
	session, err := store.Get(c)
	if err != nil {
		return errors.Unauthorized
	}

	if session.Get(AUTH_KEY) == nil {
		return errors.Unauthorized
	}

	user.ID = session.Get(USER_ID).(uint)

	if err := user.Get(); err != nil {
		return errors.Unauthorized
	}

	if err := user.Get(); err != nil {
		return NotFoundError
	}

	return nil
}

func (user *User) JSONResponse() *ResponseBody {
	return &ResponseBody{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
	}
}

func setSession(c *fiber.Ctx, user *User) error {
	session, err := store.Get(c)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	session.Set(AUTH_KEY, true)
	session.Set(USER_ID, user.ID)

	if err := session.Save(); err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}

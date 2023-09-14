package auth

import (
	"errors"

	e "github.com/guemidiborhane/explore-go/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Create(user *User) error {
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)

	if err := database.Create(&user).Error; err != nil {
		return e.Unexpected(err.Error())
	}

	return nil
}

func Get(user *User) error {
	if err := database.First(&user, user.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.EntityNotFound("No user found")
		}

		return e.Unexpected(err.Error())
	}

	return nil
}

func GetByUsername(user *User) error {
	if err := database.Where("username = ?", user.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.EntityNotFound("No user found")
		}

		return e.Unexpected(err.Error())
	}

	return nil
}

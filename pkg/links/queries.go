package links

import (
	"errors"

	e "github.com/guemidiborhane/explore-go/errors"
	"gorm.io/gorm"
)

func All(links *[]Link) error {
	if err := database.Find(&links).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.EntityNotFound("No link found")
		}

		return e.Unexpected(err.Error())
	}

	return nil
}

func Get(link *Link) error {
	if err := database.First(&link, link.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.EntityNotFound("No link found")
		}

		return e.Unexpected(err.Error())
	}

	return nil
}

func Create(link *Link) error {
	if err := database.Create(&link).Error; err != nil {
		return e.Unexpected(err.Error())
	}

	return nil
}

func Update(link *Link) error {
	if err := database.Save(&link).Error; err != nil {
		return e.Unexpected(err.Error())
	}

	return nil
}

func Destroy(link Link) error {
	if err := database.Delete(&link).Error; err != nil {
		return e.Unexpected(err.Error())
	}

	return nil
}

func GetByShort(link *Link) error {
	if err := database.Where("short = ?", link.Short).First(&link).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.EntityNotFound("No link found")
		}

		return e.Unexpected(err.Error())
	}

	return nil
}

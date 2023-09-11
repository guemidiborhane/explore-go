package queries

import (
	"errors"
	"links/models"

	"core/database"
	e "core/errors"

	"gorm.io/gorm"
)

func All(links *[]models.Link) error {
	if err := database.Database.Find(&links).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.EntityNotFound("No link found")
		}

		return e.Unexpected(err.Error())
	}

	return nil
}

func Get(link *models.Link, id uint64) error {
	if err := database.Database.First(&link, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.EntityNotFound("No link found")
		}

		return e.Unexpected(err.Error())
	}

	return nil
}

func Create(link *models.Link) error {
	if err := database.Database.Create(&link).Error; err != nil {
		return e.Unexpected(err.Error())
	}

	return nil
}

func Update(link *models.Link) error {
	if err := database.Database.Save(&link).Error; err != nil {
		return e.Unexpected(err.Error())
	}

	return nil
}

func Destroy(link models.Link) error {
	if err := database.Database.Delete(&link).Error; err != nil {
		return e.Unexpected(err.Error())
	}

	return nil
}

func GetByShort(link *models.Link, short string) error {
	if err := database.Database.Where("short = ?", short).First(&link).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.EntityNotFound("No link found")
		}

		return e.Unexpected(err.Error())
	}

	return nil
}

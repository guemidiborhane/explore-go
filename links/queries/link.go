package queries

import (
	"links/models"

	application "core/config"
)

func All() ([]models.Link, error) {
	var links []models.Link

	tx := application.Database.Find(&links)

	if tx.Error != nil {
		return []models.Link{}, tx.Error
	}

	return links, nil
}

func Get(link *models.Link, id uint64) error {
	return application.Database.First(&link, id).Error
}

func Create(link *models.Link) error {
	return application.Database.Create(&link).Error
}

func Update(link *models.Link) error {
	return application.Database.Save(&link).Error
}

func Destroy(link models.Link) error {
	return application.Database.Delete(&link).Error
}

func GetByShort(link *models.Link, short string) error {
	return application.Database.Where("short = ?", short).First(&link).Error
}

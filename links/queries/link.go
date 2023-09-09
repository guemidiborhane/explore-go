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

func Get(id uint64) (models.Link, error) {
	var link models.Link
	tx := application.Database.First(&link, id)

	if tx.Error != nil {
		return models.Link{}, tx.Error
	}

	return link, nil
}

func Create(link *models.Link) error {
	tx := application.Database.Create(&link)

	return tx.Error
}

func Update(link models.Link) error {
	tx := application.Database.Save(&link)

	return tx.Error
}

func Destroy(link models.Link) error {
	tx := application.Database.Delete(&link)

	return tx.Error
}

func GetByShort(short string) (models.Link, error) {
	var link models.Link
	tx := application.Database.Where("short = ?", short).First(&link)

	if tx.Error != nil {
		return models.Link{}, tx.Error
	}

	return link, nil
}

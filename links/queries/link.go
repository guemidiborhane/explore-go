package queries

import (
	"links/models"

	application "core/config"
)

func GetAllLinks() ([]models.Link, error) {
	var links []models.Link

	tx := application.Database.Find(&links)

	if tx.Error != nil {
		return []models.Link{}, tx.Error
	}

	return links, nil
}

func GetLink(id uint64) (models.Link, error) {
	var link models.Link
	tx := application.Database.First(&link, id)

	if tx.Error != nil {
		return models.Link{}, tx.Error
	}

	return link, nil
}

func CreateLink(link *models.Link) error {
	tx := application.Database.Create(&link)

	return tx.Error
}

func UpdateLink(link models.Link) error {
	tx := application.Database.Save(&link)

	return tx.Error
}

func DestroyLink(link models.Link) error {
	tx := application.Database.Delete(&link)

	return tx.Error
}

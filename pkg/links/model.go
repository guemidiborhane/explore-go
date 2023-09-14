package links

import (
	"gorm.io/gorm"
)

var database *gorm.DB

func setupModels() {
	database.AutoMigrate(&Link{})
}

type Link struct {
	gorm.Model
	ID    uint   `json:"id"    gorm:"primaryKey"`
	Link  string `json:"link"  gorm:"not null"        validate:"required"`
	Short string `json:"short" gorm:"unique,not null"`
}

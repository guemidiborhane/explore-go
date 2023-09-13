package links

import (
	"gorm.io/gorm"
)

var database *gorm.DB

type Link struct {
	gorm.Model
	Link  string `json:"link", gorm:"not null"`
	Short string `json:"short", gorm:"unique,not null"`
}

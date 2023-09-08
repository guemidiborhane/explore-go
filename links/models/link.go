package models

import (
	application "github.com/guemidiborhane/explore-go/config"
	"gorm.io/gorm"
)

type Link struct {
	Link  string `json:"link", gorm:"not null"`
	Short string `json:"short", gorm:"unique,not null"`
	gorm.Model
}

func Setup() {
	application.Database.AutoMigrate(&Link{})
}

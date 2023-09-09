package models

import (
	"gorm.io/gorm"
)

type Link struct {
	Link  string `json:"link", gorm:"not null"`
	Short string `json:"short", gorm:"unique,not null"`
	gorm.Model
}

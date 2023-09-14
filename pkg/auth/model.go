package auth

import "gorm.io/gorm"

var database *gorm.DB

func setupModels() {
	database.AutoMigrate(&User{})
}

type (
	User struct {
		gorm.Model
		ID       uint   `json:"id"       gorm:"primaryKey"`
		Name     string `json:"name"     validate:"required,min=5,max=20"`
		Username string `json:"username" validate:"required,min=3,max=10" gorm:"uniqueIndex"`
		Password string `json:"password" validate:"required"`
	}
)

package initializers

import (
	application "github.com/guemidiborhane/explore-go/config"
	link "github.com/guemidiborhane/explore-go/links/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := "host=127.0.0.1 user=borhane password=secret dbname=gorm port=5432 sslmode=disable TimeZone=Africa/Algiers"
	var err error
	application.Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	link.Setup()
}

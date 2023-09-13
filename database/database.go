package database

import (
	"fmt"

	"github.com/guemidiborhane/explore-go/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	Host     string = utils.GetEnv("DB_HOST", "localhost")
	Username string = utils.GetEnv("DB_USERNAME", "postgres")
	Password string = utils.GetEnv("DB_PASSWORD", "postgres")
	Name     string = utils.GetEnv("DB_NAME", "gorm")
	Port     uint64 = utils.ParseUint(utils.GetEnv("DB_PORT", "5432"), 64)
	TimeZone string = utils.GetEnv("TZ", "Africa/Algiers")
)

func Setup() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		Host, Username, Password, Name, Port, TimeZone,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
}

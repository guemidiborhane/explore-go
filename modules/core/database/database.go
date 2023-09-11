package database

import (
	"core/utils"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

var Host string = utils.GetEnv("DB_HOST", "localhost")
var Username string = utils.GetEnv("DB_USERNAME", "postgres")
var Password string = utils.GetEnv("DB_PASSWORD", "postgres")
var Name string = utils.GetEnv("DB_NAME", "gorm")
var Port uint64 = utils.ParseUint(utils.GetEnv("DB_PORT", "5432"), 64)
var TimeZone string = utils.GetEnv("TZ", "Africa/Algiers")

func Setup() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		Host, Username, Password, Name, Port, TimeZone,
	)
	var err error
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
}

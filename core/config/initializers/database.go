package initializers

import (
	application "core/config"
	"core/utils"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		utils.GetEnv("DB_HOST", "localhost"),
		utils.GetEnv("DB_USERNAME", "postgres"),
		utils.GetEnv("DB_PASSWORD", "postgres"),
		utils.GetEnv("DB_NAME", "gorm"),
		utils.ParseUint(utils.GetEnv("DB_PORT", "5432"), 64),
		utils.GetEnv("TZ", "Africa/Algiers"),
	)
	var err error
	application.Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

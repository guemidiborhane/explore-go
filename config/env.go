package config

import (
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		os.Stderr.WriteString("Error loading .env file")
	}
}

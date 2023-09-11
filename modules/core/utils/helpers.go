package utils

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ParseUint(str string, bitSize int) uint64 {
	value, err := strconv.ParseUint(str, 10, bitSize)

	if err != nil {
		panic(err)
	}

	return value
}

func HandleError(err error, c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":  "error",
		"message": err.Error(),
	})
}

func GetEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}

package utils

import (
	"math/rand"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandomShort(size uint64) string {

	str := make([]rune, size)

	for i := range str {
		str[i] = runes[rand.Intn(len(runes))]
	}

	return string(str)
}

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
		"message": err,
	})
}

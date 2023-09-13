package utils

import (
	"os"
	"strconv"

	"github.com/google/uuid"
)

func ParseUint(str string, bitSize int) uint64 {
	value, err := strconv.ParseUint(str, 10, bitSize)
	if err != nil {
		panic(err)
	}

	return value
}

func ParseInt(str string) int {
	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}

	return int(value)
}

func GetEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}

func UUIDv4() string {
	return uuid.NewString()
}

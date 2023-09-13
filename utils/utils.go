package utils

import (
	"os"
	"strconv"
)

func ParseUint(str string, bitSize int) uint64 {
	value, err := strconv.ParseUint(str, 10, bitSize)
	if err != nil {
		panic(err)
	}

	return value
}

func GetEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}

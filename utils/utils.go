package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
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

func RandomID() string {
	input := Random(64)
	// Create a SHA-512 hasher
	hasher := sha512.New()

	// Write the input string to the hasher
	hasher.Write([]byte(input))

	// Get the SHA-512 hash as a byte slice
	hashBytes := hasher.Sum(nil)

	// Convert the byte slice to a hex-encoded string
	return hex.EncodeToString(hashBytes)
}

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func Random(size uint64) string {
	str := make([]rune, size)

	for i := range str {
		str[i] = runes[rand.Intn(len(runes))]
	}

	return string(str)
}

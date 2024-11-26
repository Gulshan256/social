package env

import (
	"os"
	"strconv"
	"time"
)

func GetString(key, fallback string) string {

	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}

	return fallback
}

func GetInt(key string, fallback int) int {
	value, exists := os.LookupEnv(key)
	if exists {
		valAsInt, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return valAsInt
	}
	return fallback
}

func GetDuration(key string, fallback int) time.Duration {
	value, exists := os.LookupEnv(key)
	if exists {
		valAsInt, err := strconv.Atoi(value)
		if err != nil {
			return time.Duration(fallback)
		}
		return time.Duration(valAsInt)
	}
	return time.Duration(fallback)
}

func GetEnv() string {
	return "development"
}

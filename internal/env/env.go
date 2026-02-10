package env

import (
	"os"
	"strconv"
	"time"
)

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}

func GetInt(key string , fallback int)int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	valAInt , err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return valAInt
}

func GetDuration(key string, fallback time.Duration) time.Duration {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	d, err := time.ParseDuration(val)
	if err != nil {
		return fallback
	}

	return d
}

package configs

import (
	"os"
)

type Config struct {
	APP_PORT  string
	RedisHost string
	RedisPort string
}

func LoadConfig() *Config {
	return &Config{
		APP_PORT:  getEnv("APP_PORT", ":8081"),
		RedisHost: getEnv("GEO_REDIS_HOST", "localhost"),
		RedisPort: getEnv("GEO_REDIS_PORT", "6379"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return fallback
}

package configs

import (
	"os"
)

type Config struct {
	APP_PORT       string
	ClickHouseHost string
	ClickHousePort string
}

func LoadConfig() *Config {
	return &Config{
		APP_PORT:       getEnv("APP_PORT", ":8083"),
		ClickHouseHost: getEnv("CLICKHOUSE_HOST", "clickhouse"),
		ClickHousePort: getEnv("CLICKHOUSE_PORT", "8123"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return fallback
}

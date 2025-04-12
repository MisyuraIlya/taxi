package configs

import "os"

type Config struct {
	APP_PORT      string
	GeoServiceURL string
}

func LoadConfig() *Config {
	return &Config{
		APP_PORT:      getEnv("APP_PORT", ":9090"),
		GeoServiceURL: getEnv("GEO_SERVICE_URL", "geo-service:8081"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

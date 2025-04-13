package configs

import "os"

type Config struct {
	APP_PORT          string
	GeoServiceURL     string
	MatchingRedisHost string
	MatchingRedisPort string
}

func LoadConfig() *Config {
	return &Config{
		APP_PORT:          getEnv("APP_PORT", ":9090"),
		GeoServiceURL:     getEnv("GEO_SERVICE_URL", "geo-service:8081"),
		MatchingRedisHost: getEnv("MATCHING_REDIS_HOST", ":redis-matching"),
		MatchingRedisPort: getEnv("MATCHIG_REDIS_PORT", ":6379"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

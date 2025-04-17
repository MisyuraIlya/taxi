package configs

import "os"

type Config struct {
	AppPort        string
	GeoServiceAddr string
	RideRedisHost  string
	RideRedisPort  string
}

func LoadConfig() *Config {
	return &Config{
		AppPort:        getEnv("APP_PORT", "8084"),
		GeoServiceAddr: getEnv("GEO_SERVICE_ADDR", "geo-service:8081"),
		RideRedisHost:  getEnv("RIDE_REDIS_HOST", "redis-ride"),
		RideRedisPort:  getEnv("RIDE_REDIS_PORT", "6379"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

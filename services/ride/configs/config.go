package configs

import "os"

type Config struct {
	AppPort             string
	NotificationService string
}

func LoadConfig() *Config {
	return &Config{
		AppPort:             getEnv("APP_PORT", "8084"),
		NotificationService: getEnv("NOTIFICAION_SERVICE", "notification-service:8082"),
	}
}

func getEnv(key string, fallback string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return fallback
}

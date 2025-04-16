package redis

import (
	"ride-service/configs"

	"github.com/redis/go-redis/v9"
)

func NewRedisPool(cfg *configs.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.RideRedisHost + ":" + cfg.RideRedisPort,
		Password: "",
		DB:       0,
	})
}

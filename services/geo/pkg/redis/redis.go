package redis

import (
	"geo-service/configs"

	"github.com/redis/go-redis/v9"
)

func NewRedisPool(cfg *configs.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
		Password: "",
		DB:       0,
	})
}

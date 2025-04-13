package redis

import (
	"matching-service/configs"

	"github.com/redis/go-redis/v9"
)

func NewRedisPool(cfg *configs.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.MatchingRedisHost + ":" + cfg.MatchingRedisPort,
		Password: "",
		DB:       0,
	})
}

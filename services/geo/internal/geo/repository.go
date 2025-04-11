package geo

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Repository interface {
	UpdateLocation(driverId string, latitude string, longitude string) error
	GetLocation(driverId string) (string, string, error)
}

type repository struct {
	redis *redis.Client
}

func NewRepository(redis *redis.Client) Repository {
	return &repository{
		redis: redis,
	}
}

func (r *repository) UpdateLocation(driverId string, latitude string, longitude string) error {
	fmt.Println("Updating location in Redis")
	return nil
}

func (r *repository) GetLocation(driverId string) (string, string, error) {
	fmt.Println("Getting location from Redis")
	return "0", "0", nil
}

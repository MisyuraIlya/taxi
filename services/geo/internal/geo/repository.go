package geo

import (
	"context"
	"fmt"
	"geo-service/proto"
	"strconv"

	"github.com/mmcloughlin/geohash"
	"github.com/redis/go-redis/v9"
)

type Repository interface {
	UpdateLocation(ctx context.Context, driverId string, latitude string, longitude string, status string) error
	GetLocation(ctx context.Context, driverId string) (string, string, error)
	FindDrivers(ctx context.Context, lat, lon float64, radius float64, limit uint32, status string) ([]*proto.Driver, error)
}

type repository struct {
	redis *redis.Client
}

func NewRepository(redis *redis.Client) Repository {
	return &repository{
		redis: redis,
	}
}

func (r *repository) UpdateLocation(ctx context.Context, driverId string, latitude string, longitude string, status string) error {
	const locationKey = "drivers:locations"
	const statusKey = "drivers:status"

	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return fmt.Errorf("invalid latitude: %v", err)
	}
	lon, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return fmt.Errorf("invalid longitude: %v", err)
	}

	intHash := int64(geohash.EncodeInt(lat, lon))

	// Update Geo-location
	if _, err = r.redis.GeoAdd(ctx, locationKey, &redis.GeoLocation{
		Name:      driverId,
		Latitude:  lat,
		Longitude: lon,
		GeoHash:   intHash,
	}).Result(); err != nil {
		return fmt.Errorf("failed to update location: %v", err)
	}

	// Update Driver Status
	if status != "active" && status != "busy" {
		return fmt.Errorf("invalid status: %s", status)
	}

	if err = r.redis.HSet(ctx, statusKey, driverId, status).Err(); err != nil {
		return fmt.Errorf("failed to update status: %v", err)
	}

	fmt.Printf("Updated location/status for driver %s: lat=%f, lon=%f, status=%s\n", driverId, lat, lon, status)
	return nil
}

func (r *repository) GetLocation(ctx context.Context, driverId string) (string, string, error) {
	const key = "drivers:locations"
	positions, err := r.redis.GeoPos(ctx, key, driverId).Result()
	if err != nil {
		return "", "", fmt.Errorf("failed to get location: %v", err)
	}
	if len(positions) == 0 || positions[0] == nil {
		return "", "", fmt.Errorf("driver %s not found", driverId)
	}

	latitude := fmt.Sprintf("%f", positions[0].Latitude)
	longitude := fmt.Sprintf("%f", positions[0].Longitude)
	return latitude, longitude, nil
}

func (r *repository) FindDrivers(ctx context.Context, lat, lon float64, radius float64, limit uint32, status string) ([]*proto.Driver, error) {
	const key = "drivers:locations"
	const statusKey = "drivers:status"

	query := &redis.GeoRadiusQuery{
		Radius:      radius,
		Unit:        "m",
		WithCoord:   true,
		WithGeoHash: true,
		Sort:        "ASC",
		Count:       int(limit),
	}

	results, err := r.redis.GeoRadius(ctx, key, lon, lat, query).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to find drivers: %v", err)
	}

	var drivers []*proto.Driver
	for _, loc := range results {
		driverStatus, err := r.redis.HGet(ctx, statusKey, loc.Name).Result()
		if err != nil {
			continue
		}

		if driverStatus == status {
			drivers = append(drivers, &proto.Driver{
				DriverId:  loc.Name,
				Latitude:  loc.Latitude,
				Longitude: loc.Longitude,
				Geohash:   uint64(loc.GeoHash),
			})

			if len(drivers) >= int(limit) {
				break
			}
		}
	}

	fmt.Printf("Found %d drivers with status '%s' within %.2f meters of (%f, %f)\n", len(drivers), status, radius, lat, lon)

	return drivers, nil
}

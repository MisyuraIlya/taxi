package matching

import (
	"context"
	"encoding/json"
	"fmt"

	geopb "matching-service/protoGeo"

	"github.com/redis/go-redis/v9"
)

// Repository defines the methods to get matching clients.
type Repository interface {
	FindClients(ctx context.Context, latitude, longitude, radius float64, limit uint32) ([]*ClientLocation, error)
	UpdateUserStatus(ctx context.Context, dto UserMatchingStatus) error
	GetUserStatus(ctx context.Context, userID string) (*UserMatchingStatus, error)
}

type repository struct {
	geoClient geopb.GeoServiceClient
	redis     *redis.Client
}

func NewRepository(geoClient geopb.GeoServiceClient, redis *redis.Client) Repository {
	return &repository{geoClient: geoClient, redis: redis}
}

func (r *repository) FindClients(ctx context.Context, latitude, longitude, radius float64, limit uint32) ([]*ClientLocation, error) {
	req := &geopb.FindDriversRequest{
		Latitude:  latitude,
		Longitude: longitude,
		Radius:    radius,
		Limit:     limit,
		Status:    "active",
	}
	resp, err := r.geoClient.FindDrivers(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(resp.GetDrivers()) == 0 {
		return nil, fmt.Errorf("no drivers found within the given radius")
	}

	var clients []*ClientLocation
	for _, d := range resp.Drivers {
		clients = append(clients, &ClientLocation{
			UserID:    d.DriverId,
			Latitude:  d.Latitude,
			Longitude: d.Longitude,
			Geohash:   d.Geohash,
		})
	}
	return clients, nil
}

func (r *repository) UpdateUserStatus(ctx context.Context, dto UserMatchingStatus) error {
	key := fmt.Sprintf("user_status:%s", dto.UserID)

	data, err := r.redis.Get(ctx, key).Result()
	var updatedStatus UserMatchingStatus

	if err == nil {
		if err := json.Unmarshal([]byte(data), &updatedStatus); err != nil {
			updatedStatus = dto
		} else {
			updatedStatus.DriverID = dto.DriverID
			updatedStatus.Status = dto.Status
			updatedStatus.CreatedAt = dto.CreatedAt
			updatedStatus.ClosedAt = dto.ClosedAt
		}
	} else if err == redis.Nil {
		updatedStatus = dto
	} else {
		return err
	}

	newData, err := json.Marshal(updatedStatus)
	if err != nil {
		return err
	}

	if err := r.redis.Set(ctx, key, newData, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (r *repository) GetUserStatus(ctx context.Context, userID string) (*UserMatchingStatus, error) {
	key := fmt.Sprintf("user_status:%s", userID)
	data, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var status UserMatchingStatus
	if err := json.Unmarshal([]byte(data), &status); err != nil {
		return nil, err
	}
	return &status, nil
}

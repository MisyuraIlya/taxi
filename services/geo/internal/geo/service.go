package geo

import (
	"context"
	"geo-service/proto"
)

type Service interface {
	UpdateLocation(ctx context.Context, driverId string, latitude string, longitude string) error
	GetLocation(ctx context.Context, driverId string) (string, string, error)
	FindDrivers(ctx context.Context, lat, lon float64, radius float64, limit uint32) ([]*proto.Driver, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) UpdateLocation(ctx context.Context, driverId string, latitude string, longitude string) error {
	return s.repository.UpdateLocation(ctx, driverId, latitude, longitude)
}

func (s *service) GetLocation(ctx context.Context, driverId string) (string, string, error) {
	return s.repository.GetLocation(ctx, driverId)
}

func (s *service) FindDrivers(ctx context.Context, lat, lon float64, radius float64, limit uint32) ([]*proto.Driver, error) {
	return s.repository.FindDrivers(ctx, lat, lon, radius, limit)
}

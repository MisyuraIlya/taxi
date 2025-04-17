package ride

import (
	"context"
	"errors"
	pb "ride-service/proto"
)

type RideService interface {
	CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error)
}

type service struct {
	repository RideRepository
}

func NewRideService(repo RideRepository) RideService {
	return &service{repository: repo}
}

func (s *service) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	if req.UserId == "" {
		return nil, errors.New("user_id is required")
	}
	if req.DriverId == "" {
		return nil, errors.New("driver_id is required")
	}
	resp, err := s.repository.CreateOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

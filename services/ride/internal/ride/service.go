package ride

import (
	"context"
	"fmt"

	"ride-service/proto"
	geoPb "ride-service/proto"

	"google.golang.org/grpc"
)

type RideService interface {
	CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error)
}

type service struct {
	repo           RideRepository
	geoClient      geoPb.GeoServiceClient
	notifyEndpoint string
}

func NewRideService(
	repo RideRepository,
	geoServiceAddr string,
	notifyEndpoint string,
) (RideService, error) {
	conn, err := grpc.Dial(geoServiceAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("dial geo service %s: %w", geoServiceAddr, err)
	}

	return &service{
		repo:           repo,
		geoClient:      geoPb.NewGeoServiceClient(conn),
		notifyEndpoint: notifyEndpoint,
	}, nil
}

func (s *service) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	// 1) persist the order
	resp, err := s.repo.CreateOrder(ctx, req)
	if err != nil {
		return nil, err
	}

	// 2) launch tracker in background (fire‑and‑forget)

	go StartTracking(
		ctx,                 // parent context: will stop when ctx is canceled
		s.geoClient,         // gRPC client to GeoService
		s.notifyEndpoint,    // e.g. "http://localhost:8082/notify/clients"
		req.UserId,          // client ID
		req.DriverId,        // driver ID
		req.PickupLatitude,  // pickup lat
		req.PickupLongitude, // pickup lon
	)

	return resp, nil
}

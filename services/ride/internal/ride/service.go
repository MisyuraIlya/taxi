package ride

import (
	"context"
	"fmt"

	geo "ride-service/geoProto"
	pb "ride-service/proto"

	"google.golang.org/grpc"
)

type RideService interface {
	CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error)
}

type service struct {
	repo           RideRepository
	geoClient      geo.GeoServiceClient
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
		geoClient:      geo.NewGeoServiceClient(conn),
		notifyEndpoint: notifyEndpoint,
	}, nil
}

func (s *service) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	fmt.Println("here3")
	fmt.Println("driverId is:" + req.DriverId)

	resp, err := s.repo.CreateOrder(ctx, req)
	if err != nil {
		return nil, err
	}

	go StartTracking(
		context.Background(),
		s.geoClient,
		s.notifyEndpoint,
		req.UserId,
		req.DriverId,
		req.PickupLatitude,
		req.PickupLongitude,
	)

	return resp, nil
}

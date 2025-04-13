package geo

import (
	"context"
	"geo-service/proto"
)

type GRPCServer struct {
	proto.UnimplementedGeoServiceServer
	service Service
}

func NewGRPCServer(service Service) *GRPCServer {
	return &GRPCServer{
		service: service,
	}
}

func (s *GRPCServer) UpdateLocation(ctx context.Context, req *proto.UpdateLocationRequest) (*proto.UpdateLocationResponse, error) {
	err := s.service.UpdateLocation(ctx, req.DriverId, req.Latitude, req.Longitude, req.Status)
	if err != nil {
		return nil, err
	}
	return &proto.UpdateLocationResponse{Message: "Location updated successfully"}, nil
}

func (s *GRPCServer) GetLocation(ctx context.Context, req *proto.GetLocationRequest) (*proto.GetLocationResponse, error) {
	lat, lon, err := s.service.GetLocation(ctx, req.DriverId)
	if err != nil {
		return nil, err
	}
	return &proto.GetLocationResponse{
		Latitude:  lat,
		Longitude: lon,
	}, nil
}

func (s *GRPCServer) FindDrivers(ctx context.Context, req *proto.FindDriversRequest) (*proto.FindDriversResponse, error) {
	drivers, err := s.service.FindDrivers(ctx, req.Latitude, req.Longitude, req.Radius, req.Limit, req.Status)
	if err != nil {
		return nil, err
	}

	return &proto.FindDriversResponse{
		Drivers: drivers,
	}, nil
}

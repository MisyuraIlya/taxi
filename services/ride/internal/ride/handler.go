// internal/ride/handler.go
package ride

import (
	"context"
	"strconv"

	"google.golang.org/grpc"
	"ride-service/proto"
	geoPb "ride-service/proto"
)

type handler struct {
	UnimplementedRideServiceServer
	service        RideService
	geoServiceAddr string
}

// NewRideHandler accepts the GeoService address
func NewRideHandler(service RideService, geoAddr string) RideServiceServer {
	return &handler{
		service:        service,
		geoServiceAddr: geoAddr,
	}
}

func (h *handler) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	// 1) Create the order in Redis
	resp, err := h.service.CreateOrder(ctx, req)
	if err != nil {
		return nil, err
	}

	// 2) Start tracking for notifications
	//    Build a short‐lived gRPC connection to GeoService
	geoConn, err := grpc.Dial(h.geoServiceAddr, grpc.WithInsecure())
	if err == nil {
		geoClient := geoPb.NewGeoServiceClient(geoConn)
		// We fire-and-forget the tracker; it will cancel if the parent context closes
		StartTracking(
			ctx,
			geoClient,
			"http://localhost:8082/notify/clients",
			req.UserId,   // client ID
			req.DriverId, // driver ID
			req.PickupLatitude,
			req.PickupLongitude,
		)
	}

	return resp, nil
}

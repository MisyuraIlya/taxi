// internal/ride/handler.go
package ride

import (
	"context"
	"ride-service/proto"
)

type handler struct {
	proto.UnimplementedRideServiceServer
	service RideService
}

func NewRideHandler(svc RideService) proto.RideServiceServer {
	return &handler{service: svc}
}

func (h *handler) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	return h.service.CreateOrder(ctx, req)
}

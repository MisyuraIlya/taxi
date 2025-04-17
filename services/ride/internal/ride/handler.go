// internal/ride/handler.go
package ride

import (
	"context"
	"log"
	ridepb "ride-service/proto"
)

type handler struct {
	ridepb.UnimplementedRideServiceServer
	service RideService
}

func NewRideHandler(service RideService) ridepb.RideServiceServer {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateOrder(ctx context.Context, req *ridepb.CreateOrderRequest) (*ridepb.CreateOrderResponse, error) {
	response, err := h.service.CreateOrder(ctx, req)
	if err != nil {
		log.Println("CreateOrder error:", err)
		return nil, err
	}
	return response, nil
}

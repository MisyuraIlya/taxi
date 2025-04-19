package ride

import (
	"context"
	"fmt"
	pb "ride-service/proto"
)

type handler struct {
	pb.UnimplementedRideServiceServer
	service RideService
}

func NewRideHandler(svc RideService) pb.RideServiceServer {
	return &handler{service: svc}
}

func (h *handler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	fmt.Println("here2")
	return h.service.CreateOrder(ctx, req)
}

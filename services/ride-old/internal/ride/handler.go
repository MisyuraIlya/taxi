package ride

import (
	"context"
	"log"

	ride "ride-service/proto"
)

// RideHandler implements the gRPC server for RideService.
type RideHandler struct {
	rideService RideService
	ride.UnimplementedRideServiceServer
}

func NewRideHandler(service RideService) *RideHandler {
	return &RideHandler{rideService: service}
}

func (h *RideHandler) CreateOrder(ctx context.Context, req *ride.CreateOrderRequest) (*ride.CreateOrderResponse, error) {
	// 1) Log what we actually received
	log.Printf("Received CreateOrder request: %+v", req)
	// 2) Call into your service layer
	o, err := h.rideService.CreateOrder(
		ctx,
		req.UserId,
		req.PickupLatitude,
		req.PickupLongitude,
		req.DropoffLatitude,
		req.DropoffLongitude,
		req.DriverId,
	)
	if err != nil {
		log.Println("CreateOrder error:", err)
		return nil, err
	}

	// 3) Before returning, log what we're about to send
	log.Printf("CreateOrder ⮕ returning: orderId=%q  driverId=%q  status=%q",
		o.ID, o.DriverID, o.Status,
	)

	// 4) Build the response—now guaranteed to carry your driverId straight through
	return &ride.CreateOrderResponse{
		OrderId:  o.ID,
		DriverId: o.DriverID,
		Status:   o.Status,
	}, nil
}

func (h *RideHandler) UpdateDriverLocation(ctx context.Context, req *ride.UpdateDriverLocationRequest) (*ride.UpdateDriverLocationResponse, error) {
	// Forward to service
	if err := h.rideService.UpdateDriverLocation(ctx, req.DriverId, req.Latitude, req.Longitude); err != nil {
		log.Println("UpdateDriverLocation error:", err)
		return nil, err
	}

	// Echo driverId back in the response
	return &ride.UpdateDriverLocationResponse{
		Message:  "Driver location updated successfully",
		DriverId: req.DriverId,
	}, nil
}

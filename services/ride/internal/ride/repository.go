package ride

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/encoding/protojson"
	pb "ride-service/proto"
)

type RideRepository interface {
	// now returning (*pb.CreateOrderResponse, error)
	CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error)
}

type rideRepository struct {
	redis *redis.Client
}

func NewRideRepository(rdb *redis.Client) RideRepository {
	return &rideRepository{redis: rdb}
}

func (r *rideRepository) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	orderID := uuid.New().String()

	data, err := protojson.Marshal(req)

	if err != nil {
		return nil, fmt.Errorf("marshal CreateOrderRequest: %w", err)
	}

	key := fmt.Sprintf("order:%s", orderID)
	if err := r.redis.Set(ctx, key, data, 0).Err(); err != nil {
		return nil, fmt.Errorf("redis SET %s: %w", key, err)
	}

	return &pb.CreateOrderResponse{
		OrderId:  orderID,
		DriverId: req.DriverId,
		Status:   "created",
	}, nil
}

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"ride-service/configs"
	"ride-service/internal/ride"
	"ride-service/pkg/redis"
	ridepb "ride-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()

	cfg := configs.LoadConfig()

	rdb := redis.NewRedisPool(cfg)
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	repo := ride.NewRideRepository(rdb)

	svc, err := ride.NewRideService(
		repo,
		cfg.GeoServiceAddr,
		"http://localhost:8082/notify/clients",
	)
	if err != nil {
		log.Fatalf("Failed to initialize RideService: %v", err)
	}

	handler := ride.NewRideHandler(svc)
	fmt.Println("here1")
	address := fmt.Sprintf(":%s", cfg.AppPort)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
	}

	grpcServer := grpc.NewServer()
	ridepb.RegisterRideServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	log.Printf("RideService gRPC server listening on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server error: %v", err)
	}
}

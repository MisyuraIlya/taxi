// main.go
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"ride-service/configs"
	"ride-service/internal/ride"
	"ride-service/pkg/redis"
	ridepb "ride-service/proto"
)

func main() {
	ctx := context.Background()
	cfg := configs.LoadConfig()
	rdb := redis.NewRedisPool(cfg)
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	repository := ride.NewRideRepository(rdb)
	service := ride.NewRideService(repository)
	handler := ride.NewRideHandler(service)
	port := fmt.Sprintf(":%s", cfg.AppPort)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", port, err)
	}
	grpcServer := grpc.NewServer()
	ridepb.RegisterRideServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)
	log.Printf("gRPC server listening on %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}

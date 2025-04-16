// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"ride-service/configs"
	"ride-service/internal/ride"
	"ride-service/pkg/redis"
	ridepb "ride-service/proto"

	geo "ride-service/protoGeo"
)

func main() {
	ctx := context.Background()
	cfg := configs.LoadConfig()

	// Redis setup...
	rdb := redis.NewRedisPool(cfg)
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// Dial the GeoService
	geoConn, err := grpc.Dial(cfg.GeoServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to GeoService at %s: %v", cfg.GeoServiceAddr, err)
	}
	defer geoConn.Close()
	geoClient := geo.NewGeoServiceClient(geoConn)

	// Create your repository and service (passing geoClient in)
	repo := ride.NewRideRepository(rdb)
	service := ride.NewRideService(repo, geoClient)
	handler := ride.NewRideHandler(service)

	// Start gRPC server as before...
	addr := fmt.Sprintf(":%s", cfg.AppPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", addr, err)
	}
	grpcServer := grpc.NewServer()
	ridepb.RegisterRideServiceServer(grpcServer, handler)

	log.Printf("gRPC server listening on %s", addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

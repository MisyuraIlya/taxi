package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"matching-service/configs"
	"matching-service/internal/matching"
	"matching-service/pkg/redis"

	pb "matching-service/proto"
	geopb "matching-service/protoGeo"
)

func main() {
	cfg := configs.LoadConfig()
	redisClient := redis.NewRedisPool(cfg)
	geoConn, err := grpc.Dial(cfg.GeoServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Printf("Connecting to geo service at %s\n", cfg.GeoServiceURL)
	if err != nil {
		log.Fatalf("failed to connect to geo service: %v", err)
	}
	defer geoConn.Close()

	geoClient := geopb.NewGeoServiceClient(geoConn)

	repo := matching.NewRepository(geoClient)
	svc := matching.NewService(repo)
	handler := matching.NewHandler(svc)

	grpcServer := grpc.NewServer()
	pb.RegisterMatchingServiceServer(grpcServer, handler)

	lis, err := net.Listen("tcp", cfg.APP_PORT)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", cfg.APP_PORT, err)
	}
	log.Println("Matching service running on", cfg.APP_PORT)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"geo-service/configs"
	"geo-service/internal/geo"
	"geo-service/pkg/redis"
	"log"
	"net"

	"google.golang.org/grpc"

	"geo-service/proto"
)

func main() {
	config := configs.LoadConfig()
	redisClient := redis.NewRedisPool(config)
	repository := geo.NewRepository(redisClient)
	service := geo.NewService(repository)
	grpcServer := grpc.NewServer()
	grpcService := geo.NewGRPCServer(service)
	proto.RegisterGeoServiceServer(grpcServer, grpcService)
	lis, err := net.Listen("tcp", config.APP_PORT)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", config.APP_PORT, err)
	}

	log.Println("Geo service running on", config.APP_PORT)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

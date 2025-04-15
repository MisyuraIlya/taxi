package main

import (
	"fmt"
	"log"
	"net"

	"history-service/configs"
	pb "history-service/proto"

	"google.golang.org/grpc"

	hist "history-service/internal/history"
)

func main() {
	cfg := configs.LoadConfig()
	histRepo, err := hist.NewHistoryRepository(cfg)
	if err != nil {
		log.Fatalf("failed to connect to ClickHouse: %v", err)
	}
	if err := histRepo.Init(); err != nil {
		log.Fatalf("failed to initialize history repository: %v", err)
	}
	histService := hist.NewHistoryService(histRepo)
	grpcHandler := hist.NewGRPCHandler(histService)
	grpcServer := grpc.NewServer()
	pb.RegisterHistoryServiceServer(grpcServer, grpcHandler)
	lis, err := net.Listen("tcp", cfg.APP_PORT)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", cfg.APP_PORT, err)
	}
	fmt.Printf("History gRPC service running on %s\n", cfg.APP_PORT)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

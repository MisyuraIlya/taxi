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
	// Load configuration.
	cfg := configs.LoadConfig()

	// Initialize ClickHouse repository.
	histRepo, err := hist.NewHistoryRepository(cfg)
	if err != nil {
		log.Fatalf("failed to connect to ClickHouse: %v", err)
	}

	// Create the history table if it doesn't exist.
	if err := histRepo.Init(); err != nil {
		log.Fatalf("failed to initialize history repository: %v", err)
	}

	// Create history service and gRPC handler.
	histService := hist.NewHistoryService(histRepo)
	grpcHandler := hist.NewGRPCHandler(histService)

	// Start gRPC server.
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

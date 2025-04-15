package history

import (
	"context"
	"log"
	"time"

	pb "history-service/proto"
)

type GRPCHandler struct {
	pb.UnimplementedHistoryServiceServer
	historyService Service
}

func NewGRPCHandler(s Service) *GRPCHandler {
	return &GRPCHandler{
		historyService: s,
	}
}

func (h *GRPCHandler) CreateHistory(ctx context.Context, req *pb.CreateHistoryRequest) (*pb.CreateHistoryResponse, error) {
	createdAt, err := time.Parse(time.RFC3339, req.GetCreatedAt())
	if err != nil {
		return nil, err
	}
	closedAt, err := time.Parse(time.RFC3339, req.GetClosedAt())
	if err != nil {
		return nil, err
	}

	record := &HistoryRecord{
		UserID:    req.GetUserId(),
		DriverID:  req.GetDriverId(),
		CreatedAt: createdAt,
		ClosedAt:  closedAt,
		From:      req.GetFrom(),
		To:        req.GetTo(),
	}

	if err := h.historyService.SaveHistory(ctx, record); err != nil {
		log.Printf("Error saving history: %v", err)
		return nil, err
	}

	return &pb.CreateHistoryResponse{Message: "History saved successfully"}, nil
}

func (h *GRPCHandler) GetHistories(ctx context.Context, req *pb.GetHistoriesRequest) (*pb.GetHistoriesResponse, error) {
	histories, err := h.historyService.GetHistories(ctx)
	if err != nil {
		return nil, err
	}

	var records []*pb.HistoryRecord
	for _, rec := range histories {
		record := &pb.HistoryRecord{
			UserId:    rec.UserID,
			DriverId:  rec.DriverID,
			CreatedAt: rec.CreatedAt.Format(time.RFC3339),
			ClosedAt:  rec.ClosedAt.Format(time.RFC3339),
			From:      rec.From,
			To:        rec.To,
		}
		records = append(records, record)
	}

	return &pb.GetHistoriesResponse{Histories: records}, nil
}

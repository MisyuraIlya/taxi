package matching

import (
	"context"
	"fmt"
	pb "matching-service/proto"
	"strings"
	"time"
)

type Handler struct {
	pb.UnimplementedMatchingServiceServer
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) MatchClients(ctx context.Context, req *pb.MatchClientsRequest) (*pb.MatchClientsResponse, error) {
	clients, err := h.service.MatchClients(ctx, req.GetLatitude(), req.GetLongitude(), req.GetRadius(), req.GetLimit())
	if err != nil {
		return nil, err
	}

	var acceptedClient *ClientLocation
	notifyURL := "http://notification-service:8082/notify/driver"
	notificationMessage := "New ride request: Accept or Decline?"

	for _, c := range clients {
		response, err := SendNotification(c.UserID, notificationMessage, notifyURL)
		if err != nil {
			fmt.Printf("error sending notification to driver %s: %v\n", c.UserID, err)
			continue
		}

		fmt.Printf("Driver %s responded: %s\n", c.UserID, response)
		if strings.ToLower(response) == "accept" {
			acceptedClient = c
			break
		}
	}

	if acceptedClient == nil {
		return nil, fmt.Errorf("no driver accepted the ride")
	}

	pbClient := &pb.ClientLocation{
		UserId:    acceptedClient.UserID,
		Latitude:  acceptedClient.Latitude,
		Longitude: acceptedClient.Longitude,
		Geohash:   acceptedClient.Geohash,
	}

	return &pb.MatchClientsResponse{
		Clients: []*pb.ClientLocation{pbClient},
	}, nil
}

func (h *Handler) UpdateUserStatus(ctx context.Context, req *pb.UpdateUserStatusRequest) (*pb.UpdateUserStatusResponse, error) {
	status := UserMatchingStatus{
		UserID:    req.GetUserId(),
		DriverID:  req.GetDriverId(),
		Status:    Status(req.GetStatus()),
		CreatedAt: time.Now(),
	}

	if err := h.service.UpdateUserStatus(ctx, status); err != nil {
		return nil, fmt.Errorf("failed to update user status: %w", err)
	}
	return &pb.UpdateUserStatusResponse{
		Message: "User status updated successfully",
	}, nil
}

func (h *Handler) GetUserMatchingStatus(ctx context.Context, req *pb.GetUserMatchingStatusRequest) (*pb.GetUserMatchingStatusResponse, error) {
	status, err := h.service.GetUserStatus(ctx, req.GetUserId())
	if err != nil {
		return nil, fmt.Errorf("could not fetch status for user %s: %w", req.GetUserId(), err)
	}

	resp := &pb.GetUserMatchingStatusResponse{
		UserId:    status.UserID,
		DriverId:  status.DriverID,
		Status:    string(status.Status),
		CreatedAt: status.CreatedAt.Format(time.RFC3339),
	}
	if status.ClosedAt != nil {
		resp.ClosedAt = status.ClosedAt.Format(time.RFC3339)
	}
	return resp, nil
}

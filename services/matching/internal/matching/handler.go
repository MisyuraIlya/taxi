package matching

import (
	"context"
	"fmt"
	"strings"

	pb "matching-service/proto"
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

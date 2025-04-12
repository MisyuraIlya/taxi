package matching

import (
	"context"

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

	var pbClients []*pb.ClientLocation
	for _, c := range clients {
		pbClients = append(pbClients, &pb.ClientLocation{
			UserId:    c.UserID,
			Latitude:  c.Latitude,
			Longitude: c.Longitude,
			Geohash:   c.Geohash,
		})
	}

	return &pb.MatchClientsResponse{
		Clients: pbClients,
	}, nil
}

package matching

import "context"

type Service interface {
	MatchClients(ctx context.Context, latitude, longitude, radius float64, limit uint32) ([]*ClientLocation, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) MatchClients(ctx context.Context, latitude, longitude, radius float64, limit uint32) ([]*ClientLocation, error) {
	return s.repository.FindClients(ctx, latitude, longitude, radius, limit)
}

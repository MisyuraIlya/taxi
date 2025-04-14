package history

import (
	"context"
)

// Service defines business logic for handling history records.
type Service interface {
	SaveHistory(ctx context.Context, record *HistoryRecord) error
	GetHistories(ctx context.Context) ([]HistoryRecord, error)
}

type historyService struct {
	repo Repository
}

// NewHistoryService creates a new history service.
func NewHistoryService(repo Repository) Service {
	return &historyService{repo: repo}
}

// SaveHistory saves a history record.
func (s *historyService) SaveHistory(ctx context.Context, record *HistoryRecord) error {
	return s.repo.SaveHistory(ctx, record)
}

// GetHistories retrieves all history records.
func (s *historyService) GetHistories(ctx context.Context) ([]HistoryRecord, error) {
	return s.repo.GetHistories(ctx)
}

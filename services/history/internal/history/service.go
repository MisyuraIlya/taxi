package history

import (
	"context"
)

type Service interface {
	SaveHistory(ctx context.Context, record *HistoryRecord) error
	GetHistories(ctx context.Context) ([]HistoryRecord, error)
}

type historyService struct {
	repo Repository
}

func NewHistoryService(repo Repository) Service {
	return &historyService{repo: repo}
}

func (s *historyService) SaveHistory(ctx context.Context, record *HistoryRecord) error {
	return s.repo.SaveHistory(ctx, record)
}

func (s *historyService) GetHistories(ctx context.Context) ([]HistoryRecord, error) {
	return s.repo.GetHistories(ctx)
}

package history

import (
	"context"
	"fmt"

	"history-service/configs"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type Repository interface {
	Init() error
	SaveHistory(ctx context.Context, record *HistoryRecord) error
	GetHistories(ctx context.Context) ([]HistoryRecord, error)
}

type repository struct {
	conn clickhouse.Conn
}

func NewHistoryRepository(cfg *configs.Config) (Repository, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", cfg.ClickHouseHost, cfg.ClickHousePort)},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "mysecretpassword",
		},
	})
	if err != nil {
		return nil, err
	}

	return &repository{conn: conn}, nil
}

func (r *repository) Init() error {
	ctx := context.Background()
	query := `
    CREATE TABLE IF NOT EXISTS history (
        userId String,
        driverId String,
        createdAt DateTime,
        closedAt DateTime,
        ` + "`from`" + ` String,
        ` + "`to`" + ` String
    ) ENGINE = MergeTree()
    ORDER BY createdAt
    `
	return r.conn.Exec(ctx, query)
}

func (r *repository) SaveHistory(ctx context.Context, record *HistoryRecord) error {
	query := `
    INSERT INTO history (userId, driverId, createdAt, closedAt, ` + "`from`" + `, ` + "`to`" + `) 
    VALUES (?, ?, ?, ?, ?, ?)
    `
	return r.conn.Exec(ctx, query,
		record.UserID,
		record.DriverID,
		record.CreatedAt,
		record.ClosedAt,
		record.From,
		record.To,
	)
}

func (r *repository) GetHistories(ctx context.Context) ([]HistoryRecord, error) {
	query := `
    SELECT userId, driverId, createdAt, closedAt, ` + "`from`" + `, ` + "`to`" + `
    FROM history
    `
	var histories []HistoryRecord
	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rec HistoryRecord
		if err := rows.Scan(
			&rec.UserID,
			&rec.DriverID,
			&rec.CreatedAt,
			&rec.ClosedAt,
			&rec.From,
			&rec.To,
		); err != nil {
			return nil, err
		}
		histories = append(histories, rec)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return histories, nil
}

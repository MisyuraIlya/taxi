package history

import "time"

type HistoryRecord struct {
	UserID    string    `json:"userId"`
	DriverID  string    `json:"driverId"`
	CreatedAt time.Time `json:"createdAt"`
	ClosedAt  time.Time `json:"closedAt"`
	From      string    `json:"from"`
	To        string    `json:"to"`
}

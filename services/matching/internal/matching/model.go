package matching

import "time"

type ClientLocation struct {
	UserID    string  `json:"user_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Geohash   uint64  `json:"geohash"`
}

type NotificationPayload struct {
	DriverID string `json:"driver_id"`
	Message  string `json:"message"`
}

type Status string

const (
	Waiting Status = "waiting"
	Ride    Status = "ride"
	Active  Status = "active"
)

type UserMatchingStatus struct {
	UserID    string     `json:"user_id"`
	DriverID  string     `json:"driver_id"`
	Status    Status     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	ClosedAt  *time.Time `json:"closed_at,omitempty"`
}

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type PayloadMessage struct {
	ClientID string  `json:"driver_id"`
	Message  Message `json:"message"`
}

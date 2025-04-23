package matching

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Service interface {
	MatchClients(ctx context.Context, latitude, longitude, radius float64, limit uint32) ([]*ClientLocation, error)
	UpdateUserStatus(ctx context.Context, dto UserMatchingStatus) error
	GetUserStatus(ctx context.Context, userID string) (*UserMatchingStatus, error)
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

func (s *service) UpdateUserStatus(ctx context.Context, dto UserMatchingStatus) error {
	return s.repository.UpdateUserStatus(ctx, dto)
}

func (s *service) GetUserStatus(ctx context.Context, userID string) (*UserMatchingStatus, error) {
	return s.repository.GetUserStatus(ctx, userID)
}

func SendNotification(driverID, message, notifyURL string) (string, error) {
	payload := PayloadMessage{
		ClientID: driverID,
		Message: Message{
			Type: "rideRequest",
			Data: message,
		},
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	client := http.Client{Timeout: 6 * time.Second}
	resp, err := client.Post(notifyURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("notification service error: %s", string(bodyBytes))
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	resultStr := string(bodyBytes)
	parts := strings.Split(resultStr, ":")
	if len(parts) < 2 {
		return "", fmt.Errorf("unexpected response format: %s", resultStr)
	}

	responsePart := strings.TrimSpace(parts[len(parts)-1])
	return responsePart, nil
}

package ride

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RideRepository interface {
	CreateOrder(ctx context.Context, order *Order) error
	GetOrderByID(ctx context.Context, orderID string) (*Order, error)
	UpdateOrder(ctx context.Context, order *Order) error

	UpdateDriverLocation(ctx context.Context, driver *Driver) error
	GetDriverByID(ctx context.Context, driverID string) (*Driver, error)
}

type rideRepository struct {
	rdb *redis.Client
}

// NewRideRepository returns a new RideRepository using Redis
func NewRideRepository(rdb *redis.Client) RideRepository {
	return &rideRepository{rdb: rdb}
}

func (r *rideRepository) CreateOrder(ctx context.Context, order *Order) error {
	// Using Redis Hash or JSON storing strategy
	key := fmt.Sprintf("order:%s", order.ID)
	err := r.rdb.HSet(ctx, key, map[string]interface{}{
		"ID":               order.ID,
		"UserID":           order.UserID,
		"PickupLatitude":   order.PickupLatitude,
		"PickupLongitude":  order.PickupLongitude,
		"DropoffLatitude":  order.DropoffLatitude,
		"DropoffLongitude": order.DropoffLongitude,
		"Status":           order.Status,
		"DriverID":         order.DriverID,
	}).Err()
	return err
}

func (r *rideRepository) GetOrderByID(ctx context.Context, orderID string) (*Order, error) {
	key := fmt.Sprintf("order:%s", orderID)
	data, err := r.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("order %s not found", orderID)
	}

	order := &Order{
		ID:       data["ID"],
		UserID:   data["UserID"],
		Status:   data["Status"],
		DriverID: data["DriverID"],
	}
	// parse floats
	fmt.Sscanf(data["PickupLatitude"], "%f", &order.PickupLatitude)
	fmt.Sscanf(data["PickupLongitude"], "%f", &order.PickupLongitude)
	fmt.Sscanf(data["DropoffLatitude"], "%f", &order.DropoffLatitude)
	fmt.Sscanf(data["DropoffLongitude"], "%f", &order.DropoffLongitude)

	return order, nil
}

func (r *rideRepository) UpdateOrder(ctx context.Context, order *Order) error {
	key := fmt.Sprintf("order:%s", order.ID)
	err := r.rdb.HSet(ctx, key, map[string]interface{}{
		"Status":   order.Status,
		"DriverID": order.DriverID,
	}).Err()
	return err
}

// ----- Driver -----

func (r *rideRepository) UpdateDriverLocation(ctx context.Context, driver *Driver) error {
	key := fmt.Sprintf("driver:%s", driver.ID)
	err := r.rdb.HSet(ctx, key, map[string]interface{}{
		"ID":        driver.ID,
		"Latitude":  driver.Latitude,
		"Longitude": driver.Longitude,
		"Status":    driver.Status,
	}).Err()
	return err
}

func (r *rideRepository) GetDriverByID(ctx context.Context, driverID string) (*Driver, error) {
	key := fmt.Sprintf("driver:%s", driverID)
	data, err := r.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("driver %s not found", driverID)
	}
	driver := &Driver{
		ID:     data["ID"],
		Status: data["Status"],
	}
	fmt.Sscanf(data["Latitude"], "%f", &driver.Latitude)
	fmt.Sscanf(data["Longitude"], "%f", &driver.Longitude)
	return driver, nil
}

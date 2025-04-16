package ride

import (
	"context"
	"fmt"
	geo "ride-service/protoGeo"

	"github.com/google/uuid"
)

// RideService defines business logic for rides.
type RideService interface {
	CreateOrder(ctx context.Context, userID string,
		pickupLat, pickupLong, dropoffLat, dropoffLong float64,
		driverID string,
	) (*Order, error)
	UpdateDriverLocation(ctx context.Context, driverID string, lat, long float64) error
}

// rideService implements RideService.
type rideService struct {
	repo      RideRepository
	geoClient geo.GeoServiceClient
}

// NewRideService constructs a new RideService.
func NewRideService(repo RideRepository, geoClient geo.GeoServiceClient) RideService {
	return &rideService{repo: repo, geoClient: geoClient}
}

// CreateOrder persists a new ride order using the provided driverID
// and notifies the GeoService that the driver is now on a ride.
func (s *rideService) CreateOrder(
	ctx context.Context,
	userID string,
	pickupLat,
	pickupLong,
	dropoffLat,
	dropoffLong float64,
	driverID string,
) (*Order, error) {
	// build the Order with DRIVER_ASSIGNED status and provided driverID
	order := &Order{
		ID:               uuid.NewString(),
		UserID:           userID,
		PickupLatitude:   pickupLat,
		PickupLongitude:  pickupLong,
		DropoffLatitude:  dropoffLat,
		DropoffLongitude: dropoffLong,
		Status:           "DRIVER_ASSIGNED",
		DriverID:         driverID,
	}

	// Persist the order in Redis
	if err := s.repo.CreateOrder(ctx, order); err != nil {
		return nil, fmt.Errorf("repo.CreateOrder: %w", err)
	}

	// Notify GeoService that the driver is now on a ride
	upd := &geo.UpdateLocationRequest{
		DriverId:  driverID,
		Latitude:  fmt.Sprintf("%f", pickupLat),
		Longitude: fmt.Sprintf("%f", pickupLong),
		Status:    "ON_RIDE",
	}
	if _, err := s.geoClient.UpdateLocation(ctx, upd); err != nil {
		// log error but do not fail order creation
		fmt.Printf("warning: geo.UpdateLocation failed: %v\n", err)
	}

	return order, nil
}

// UpdateDriverLocation forwards driver location updates to GeoService.
func (s *rideService) UpdateDriverLocation(
	ctx context.Context,
	driverID string,
	lat, long float64,
) error {
	req := &geo.UpdateLocationRequest{
		DriverId:  driverID,
		Latitude:  fmt.Sprintf("%f", lat),
		Longitude: fmt.Sprintf("%f", long),
		Status:    "AVAILABLE",
	}
	if _, err := s.geoClient.UpdateLocation(ctx, req); err != nil {
		return fmt.Errorf("geo.UpdateLocation failed: %w", err)
	}
	return nil
}

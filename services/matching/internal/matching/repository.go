package matching

import (
	"context"
	"fmt"

	geopb "matching-service/protoGeo"
)

// Repository defines the methods to get matching clients.
type Repository interface {
	FindClients(ctx context.Context, latitude, longitude, radius float64, limit uint32) ([]*ClientLocation, error)
}

type repository struct {
	geoClient geopb.GeoServiceClient
}

func NewRepository(geoClient geopb.GeoServiceClient) Repository {
	return &repository{geoClient: geoClient}
}

func (r *repository) FindClients(ctx context.Context, latitude, longitude, radius float64, limit uint32) ([]*ClientLocation, error) {
	req := &geopb.FindDriversRequest{
		Latitude:  latitude,
		Longitude: longitude,
		Radius:    radius,
		Limit:     limit,
	}
	fmt.Println("req", req)
	resp, err := r.geoClient.FindDrivers(ctx, req)
	fmt.Println("resp", resp)
	if err != nil {
		return nil, err
	}

	var clients []*ClientLocation
	for _, d := range resp.Drivers {
		clients = append(clients, &ClientLocation{
			UserID:    d.DriverId,
			Latitude:  d.Latitude,
			Longitude: d.Longitude,
			Geohash:   d.Geohash,
		})
	}
	return clients, nil
}

package main

import (
	"ride-service/configs"
	"ride-service/internal/ride"
)

func main() {
	cfg := configs.LoadConfig()
	svc := ride.NewRideService()

}

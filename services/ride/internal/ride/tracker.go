package ride

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	geo "ride-service/geoProto"
)

func distance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000
	toRad := func(deg float64) float64 { return deg * math.Pi / 180 }
	dLat := toRad(lat2 - lat1)
	dLon := toRad(lon2 - lon1)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(toRad(lat1))*math.Cos(toRad(lat2))*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func StartTracking(
	ctx context.Context,
	geoClient geo.GeoServiceClient,
	notifyURL,
	clientID,
	driverID string,
	pickupLat,
	pickupLon float64,
) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	fmt.Println("here5")
	comingSoonNotified := false
	arrivedNotified := false

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			locResp, err := geoClient.GetLocation(ctx, &geo.GetLocationRequest{DriverId: driverID})
			if err != nil {
				fmt.Printf("tracking: GetLocation error: %v\n", err)
				continue
			}

			dLat, _ := strconv.ParseFloat(locResp.Latitude, 64)
			dLon, _ := strconv.ParseFloat(locResp.Longitude, 64)
			dist := distance(pickupLat, pickupLon, dLat, dLon)
			fmt.Printf("driver %s: %f\n", driverID, dist)
			if dist <= 50 && !comingSoonNotified {
				sendClientNotification(notifyURL, clientID, "Driver is within 50 meters")
				comingSoonNotified = true
			}

			if dist <= 5 && !arrivedNotified {
				sendClientNotification(notifyURL, clientID, "Driver has arrived")
				arrivedNotified = true
			}
		}
	}
}

func sendClientNotification(url, clientID, message string) {
	fmt.Printf("notifying client %s: %s\n", clientID, message)
	payload := map[string]string{
		"client_id": clientID,
		"message":   message,
	}
	body, _ := json.Marshal(payload)
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	fmt.Println(resp)
	if err != nil {
		fmt.Printf("notify error: %v\n", err)
		return
	}
	resp.Body.Close()
}

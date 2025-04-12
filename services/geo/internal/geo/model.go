package geo

type Geo struct {
	DriverId  string `json:"driver_id"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	GeoHash   string `json:"geohash"`
}

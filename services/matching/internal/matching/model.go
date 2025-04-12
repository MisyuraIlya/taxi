package matching

type ClientLocation struct {
	UserID    string  `json:"user_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Geohash   uint64  `json:"geohash"`
}

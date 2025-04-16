package ride

// Order represents a ride order
type Order struct {
	ID               string
	UserID           string
	PickupLatitude   float64
	PickupLongitude  float64
	DropoffLatitude  float64
	DropoffLongitude float64
	Status           string // e.g. "REQUESTED", "DRIVER_ASSIGNED", etc.
	DriverID         string
}

// Driver represents a driver in the system
type Driver struct {
	ID        string
	Latitude  float64
	Longitude float64
	Status    string // e.g. "AVAILABLE", "ON_RIDE", "OFFLINE"
}

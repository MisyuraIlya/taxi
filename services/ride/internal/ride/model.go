package ride

// CreateOrderRequest represents a request to create a new ride order.
type CreateOrderRequest struct {
	UserID           string  `json:"user_id"`
	DriverID         string  `json:"driver_id"`
	PickupLatitude   float64 `json:"pickup_latitude"`
	PickupLongitude  float64 `json:"pickup_longitude"`
	DropoffLatitude  float64 `json:"dropoff_latitude"`
	DropoffLongitude float64 `json:"dropoff_longitude"`
}

type CreateOrderResponse struct {
	OrderID  string `json:"order_id"`
	DriverID string `json:"driver_id"`
	Status   string `json:"status"`
}

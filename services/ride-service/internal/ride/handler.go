package ride

import "context"

type Handler interface {
	NearByNotification(ctx context.Context, req any)
	DriverWaitingNotification(ctx context.Context, req any)
}

type handler struct {
	service service
}

func NewRideHandler(service service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) NearByNotification(ctx context.Context, req any) {

}

func (h *handler) DriverWaitingNotification(ctx context.Context, req any) {

}

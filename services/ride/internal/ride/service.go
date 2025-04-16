package ride

import "context"

type Service interface {
	NearByNotification(ctx context.Context)
	DriverWaitingNotification(ctx context.Context)
}

type service struct {
}

func NewRideService() Service {
	return &service{}
}

func (s *service) NearByNotification(ctx context.Context) {

}

func (s *service) DriverWaitingNotification(ctx context.Context) {

}

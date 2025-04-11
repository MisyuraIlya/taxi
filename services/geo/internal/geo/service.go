package geo

type Service interface {
	UpdateLocation(driverId string, latitude string, longitude string) error
	GetLocation(driverId string) (string, string, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) UpdateLocation(driverId string, latitude string, longitude string) error {
	err := s.repository.UpdateLocation(driverId, latitude, longitude)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetLocation(driverId string) (string, string, error) {
	latitude, longitude, err := s.repository.GetLocation(driverId)
	if err != nil {
		return "", "", err
	}
	return latitude, longitude, nil
}

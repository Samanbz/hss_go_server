package services

import (
	"context"
	"hss/internal/models"
	"hss/internal/repositories"
)

// I know, I know. Sorry for the name.
type ServiceService struct {
	serviceRepository *repositories.ServiceRepository
}

func NewServiceService(serviceRepository *repositories.ServiceRepository) *ServiceService {
	return &ServiceService{serviceRepository: serviceRepository}
}

func (s *ServiceService) InsertService(ctx context.Context, service *models.Service) error {
	err := s.serviceRepository.InsertService(context.Background(), service)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceService) GetServiceByID(ctx context.Context, id int) (*models.Service, error) {
	service, err := s.serviceRepository.GetServiceByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (s *ServiceService) GetServicesByAddressID(ctx context.Context, addressID int) (*[]models.Service, error) {
	services, err := s.serviceRepository.GetServicesByAddressID(context.Background(), addressID)
	if err != nil {
		return nil, err
	}
	return services, nil
}

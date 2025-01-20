package services

import (
	"context"
	"hss/internal/models"
	"hss/internal/repositories"
)

type CustomerService struct {
	customerRepository *repositories.CustomerRepository
}

func NewCustomerService(customerRepository *repositories.CustomerRepository) *CustomerService {
	return &CustomerService{customerRepository: customerRepository}
}

func (s *CustomerService) InsertCustomer(ctx context.Context, customer *models.Customer) error {
	err := s.customerRepository.Create(context.Background(), customer)
	if err != nil {
		return err
	}
	return nil
}

func (s *CustomerService) GetCustomerByID(ctx context.Context, id int) (*models.Customer, error) {
	customer, err := s.customerRepository.GetByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

package services

import (
	"context"
	"hss/internal/api/repositories"
	"hss/internal/models"
)

type CustomerService struct {
	customerRepository *repositories.CustomerRepository
}

func NewCustomerService(customerRepository *repositories.CustomerRepository) *CustomerService {
	return &CustomerService{customerRepository: customerRepository}
}

func (s *CustomerService) InsertCustomer(ctx context.Context, customer *models.Customer) error {
	err := s.customerRepository.InsertCustomer(context.Background(), customer)
	if err != nil {
		return err
	}
	return nil
}

func (s *CustomerService) GetCustomerByID(ctx context.Context, id int) (*models.Customer, error) {
	customer, err := s.customerRepository.GetCustomerByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

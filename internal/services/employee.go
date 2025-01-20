package services

import (
	"context"
	"hss/internal/models"
	"hss/internal/repositories"
)

type EmployeeService struct {
	employeeRepository *repositories.EmployeeRepository
}

func NewEmployeeService(employeeRepository *repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{employeeRepository: employeeRepository}
}

func (s *EmployeeService) InsertEmployee(ctx context.Context, employee *models.Employee) error {
	err := s.employeeRepository.Create(context.Background(), employee)
	if err != nil {
		return err
	}
	return nil
}

func (s *EmployeeService) GetEmployeeByID(ctx context.Context, id int) (*models.Employee, error) {
	employee, err := s.employeeRepository.GetByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

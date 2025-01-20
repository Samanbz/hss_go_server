package mocks

import (
	"context"
	"hss/internal/models"
	"hss/internal/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeMockGroup struct {
	mockObjects []*models.Employee
}

func NewEmployeeMockGroup(mockObjects ...*models.Employee) EmployeeMockGroup {
	return EmployeeMockGroup{
		mockObjects: mockObjects,
	}
}

func (m EmployeeMockGroup) LoadSelf(ctx context.Context, pool *pgxpool.Pool) error {
	employeeRepository := repositories.NewEmployeeRepository(pool)

	for _, employee := range m.mockObjects {
		err := employeeRepository.Create(ctx, employee)
		if err != nil {
			return err
		}
	}
	return nil
}

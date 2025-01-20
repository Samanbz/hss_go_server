package repositories

import (
	"context"
	"hss/internal/models"
)

type Repository[T models.BaseModel] interface {
	Create(context.Context, *T) error
	GetByID(context.Context, int) (*T, error)
	UpdateByID(context.Context, int, *T) error
	DeleteByID(context.Context, int) (*T, error)
}

type Repositories struct {
	CompanyRepository     *CompanyRepository
	AddressRepository     *AddressRepository
	EmployeeRepository    *EmployeeRepository
	AppointmentRepository *AppointmentRepository
	ServiceRepository     *ServiceRepository
	CustomerRepository    *CustomerRepository
	ProductRepository     *ProductRepository
	AuthRepository        *AuthRepository
}

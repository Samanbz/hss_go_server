package singleton

import (
	"hss/internal/api/handlers"
	"hss/internal/api/repositories"
	"hss/internal/api/services"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitCompanySingletons(pool *pgxpool.Pool) (*handlers.CompanyHandler, error) {
	//TODO add error cases
	CompanyRepository := repositories.NewCompanyRepository(pool)
	CompanyService := services.NewCompanyService(CompanyRepository)
	CompanyHandler := handlers.NewCompanyHandler(CompanyService)

	return &CompanyHandler, nil
}

func InitAddressSingletons(pool *pgxpool.Pool) (*handlers.AddressHandler, error) {
	//TODO add error cases
	AddressRepository := repositories.NewAddressRepository(pool)
	AddressService := services.NewAddressService(AddressRepository)
	AddressHandler := handlers.NewAddressHandler(AddressService)

	return &AddressHandler, nil
}

func InitEmployeeSingletons(pool *pgxpool.Pool) (*handlers.EmployeeHandler, error) {
	//TODO add error cases
	EmployeeRepository := repositories.NewEmployeeRepository(pool)
	EmployeeService := services.NewEmployeeService(EmployeeRepository)
	EmployeeHandler := handlers.NewEmployeeHandler(EmployeeService)

	return &EmployeeHandler, nil
}

func InitAppointmentsSingletons(pool *pgxpool.Pool) (*handlers.AppointmentHandler, error) {
	//TODO add error cases
	AppointmentRepository := repositories.NewAppointmentRepository(pool)
	AppointmentService := services.NewAppointmentService(AppointmentRepository)
	AppointmentHandler := handlers.NewAppointmentHandler(AppointmentService)

	return &AppointmentHandler, nil
}

func InitSingletons(pool *pgxpool.Pool) (*handlers.CompanyHandler, *handlers.AddressHandler, *handlers.EmployeeHandler, *handlers.AppointmentHandler, error) {
	CompanyHandler, err := InitCompanySingletons(pool)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	AddressHandler, err := InitAddressSingletons(pool)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	EmployeeHandler, err := InitEmployeeSingletons(pool)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	AppointmentHandler, err := InitAppointmentsSingletons(pool)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return CompanyHandler, AddressHandler, EmployeeHandler, AppointmentHandler, nil
}

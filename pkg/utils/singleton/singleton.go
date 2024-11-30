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

	return CompanyHandler, nil
}

func InitAddressSingletons(pool *pgxpool.Pool) (*handlers.AddressHandler, error) {
	//TODO add error cases
	AddressRepository := repositories.NewAddressRepository(pool)
	AddressService := services.NewAddressService(AddressRepository)
	AddressHandler := handlers.NewAddressHandler(AddressService)

	return AddressHandler, nil
}

func InitEmployeeSingletons(pool *pgxpool.Pool) (*handlers.EmployeeHandler, error) {
	//TODO add error cases
	EmployeeRepository := repositories.NewEmployeeRepository(pool)
	EmployeeService := services.NewEmployeeService(EmployeeRepository)
	EmployeeHandler := handlers.NewEmployeeHandler(EmployeeService)

	return EmployeeHandler, nil
}

func InitAppointmentsSingletons(pool *pgxpool.Pool) (*handlers.AppointmentHandler, error) {
	//TODO add error cases
	AppointmentRepository := repositories.NewAppointmentRepository(pool)
	AppointmentService := services.NewAppointmentService(AppointmentRepository)
	AppointmentHandler := handlers.NewAppointmentHandler(AppointmentService)

	return AppointmentHandler, nil
}

func InitServiceSingletons(pool *pgxpool.Pool) (*handlers.ServiceHandler, error) {
	//TODO add error cases
	ServiceRepository := repositories.NewServiceRepository(pool)
	ServiceService := services.NewServiceService(ServiceRepository)
	ServiceHandler := handlers.NewServiceHandler(ServiceService)

	return ServiceHandler, nil
}

func InitCustomerSingletons(pool *pgxpool.Pool) (*handlers.CustomerHandler, error) {
	//TODO add error cases
	CustomerRepository := repositories.NewCustomerRepository(pool)
	CustomerService := services.NewCustomerService(CustomerRepository)
	CustomerHandler := handlers.NewCustomerHandler(CustomerService)

	return CustomerHandler, nil
}

func InitProductSingletons(pool *pgxpool.Pool) (*handlers.ProductHandler, error) {
	//TODO add error cases
	ProductRepository := repositories.NewProductRepository(pool)
	ProductService := services.NewProductService(ProductRepository)
	ProductHandler := handlers.NewProductHandler(ProductService)

	return ProductHandler, nil
}

func InitSingletons(pool *pgxpool.Pool) (*handlers.RequestHandlers, error) {
	CompanyHandler, err := InitCompanySingletons(pool)
	if err != nil {
		return nil, err
	}

	AddressHandler, err := InitAddressSingletons(pool)
	if err != nil {
		return nil, err
	}

	EmployeeHandler, err := InitEmployeeSingletons(pool)
	if err != nil {
		return nil, err
	}

	AppointmentHandler, err := InitAppointmentsSingletons(pool)
	if err != nil {
		return nil, err
	}

	ServiceHandler, err := InitServiceSingletons(pool)
	if err != nil {
		return nil, err
	}

	CustomerHandler, err := InitCustomerSingletons(pool)
	if err != nil {
		return nil, err
	}

	ProductHandler, err := InitProductSingletons(pool)
	if err != nil {
		return nil, err
	}

	return &handlers.RequestHandlers{
		CompanyHandler:     CompanyHandler,
		AddressHandler:     AddressHandler,
		EmployeeHandler:    EmployeeHandler,
		AppointmentHandler: AppointmentHandler,
		ServiceHandler:     ServiceHandler,
		CustomerHandler:    CustomerHandler,
		ProductHandler:     ProductHandler,
	}, nil
}

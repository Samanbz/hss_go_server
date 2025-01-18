package singleton

import (
	"hss/internal/handlers"
	"hss/internal/repositories"
	"hss/internal/services"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitCompanySingletons(pool *pgxpool.Pool) *handlers.CompanyHandler {
	CompanyRepository := repositories.NewCompanyRepository(pool)
	CompanyService := services.NewCompanyService(CompanyRepository)
	CompanyHandler := handlers.NewCompanyHandler(CompanyService)

	return CompanyHandler
}

func InitAddressSingletons(pool *pgxpool.Pool) *handlers.AddressHandler {
	AddressRepository := repositories.NewAddressRepository(pool)
	AddressService := services.NewAddressService(AddressRepository)
	AddressHandler := handlers.NewAddressHandler(AddressService)

	return AddressHandler
}

func InitEmployeeSingletons(pool *pgxpool.Pool) *handlers.EmployeeHandler {
	EmployeeRepository := repositories.NewEmployeeRepository(pool)
	EmployeeService := services.NewEmployeeService(EmployeeRepository)
	EmployeeHandler := handlers.NewEmployeeHandler(EmployeeService)

	return EmployeeHandler
}

func InitAppointmentsSingletons(pool *pgxpool.Pool) *handlers.AppointmentHandler {
	AppointmentRepository := repositories.NewAppointmentRepository(pool)
	AppointmentService := services.NewAppointmentService(AppointmentRepository)
	AppointmentHandler := handlers.NewAppointmentHandler(AppointmentService)

	return AppointmentHandler
}

func InitServiceSingletons(pool *pgxpool.Pool) *handlers.ServiceHandler {
	ServiceRepository := repositories.NewServiceRepository(pool)
	ServiceService := services.NewServiceService(ServiceRepository)
	ServiceHandler := handlers.NewServiceHandler(ServiceService)

	return ServiceHandler
}

func InitCustomerSingletons(pool *pgxpool.Pool) *handlers.CustomerHandler {
	CustomerRepository := repositories.NewCustomerRepository(pool)
	CustomerService := services.NewCustomerService(CustomerRepository)
	CustomerHandler := handlers.NewCustomerHandler(CustomerService)

	return CustomerHandler
}

func InitProductSingletons(pool *pgxpool.Pool) *handlers.ProductHandler {
	ProductRepository := repositories.NewProductRepository(pool)
	ProductService := services.NewProductService(ProductRepository)
	ProductHandler := handlers.NewProductHandler(ProductService)

	return ProductHandler
}

func InitAuthSingletons(pool *pgxpool.Pool) (*handlers.AuthHandler, error) {
	AuthRepository := repositories.NewAuthRepository(pool)
	AuthService, err := services.NewAuthService(AuthRepository)

	if err != nil {
		return nil, err
	}

	AuthHandler := handlers.NewAuthHandler(AuthService)

	return AuthHandler, nil
}

func InitSingletons(pool *pgxpool.Pool) (*handlers.RequestHandlers, error) {
	CompanyHandler := InitCompanySingletons(pool)

	AddressHandler := InitAddressSingletons(pool)

	EmployeeHandler := InitEmployeeSingletons(pool)

	AppointmentHandler := InitAppointmentsSingletons(pool)

	ServiceHandler := InitServiceSingletons(pool)

	CustomerHandler := InitCustomerSingletons(pool)

	ProductHandler := InitProductSingletons(pool)

	AuthHandler, err := InitAuthSingletons(pool)

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
		AuthHandler:        AuthHandler,
	}, nil
}

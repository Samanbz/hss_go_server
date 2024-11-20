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

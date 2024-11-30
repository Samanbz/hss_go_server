package test

import (
	"context"
	"hss/internal/api/repositories"
	"hss/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Mocks struct {
	companyMock *models.Company
}

func insertCompanyMock(ctx context.Context, pool *pgxpool.Pool, company *models.Company) (*models.Company, error) {
	companyRepository := repositories.NewCompanyRepository(pool)

	err := companyRepository.InsertCompany(ctx, company)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func LoadMocks(ctx context.Context, pool *pgxpool.Pool, mocks *Mocks) error {
	company, err := insertCompanyMock(ctx, pool, mocks.companyMock)
	if err != nil {
		return err
	}

	mocks.companyMock = company

	return nil
}

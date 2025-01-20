package mocks

import (
	"context"
	"fmt"
	"hss/internal/models"
	"hss/internal/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CompanyMockGroup struct {
	mockObject *models.Company
}

func NewCompanyMockGroup(mockObject *models.Company) CompanyMockGroup {
	return CompanyMockGroup{
		mockObject: mockObject,
	}
}

func (m CompanyMockGroup) LoadSelf(ctx context.Context, pool *pgxpool.Pool) error {
	companyRepository := repositories.NewCompanyRepository(pool)

	err := companyRepository.Create(ctx, m.mockObject)
	if err != nil {
		return fmt.Errorf("failed to load company mock: %w", err)
	}
	return nil
}

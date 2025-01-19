package mocks

import (
	"context"
	"fmt"
	"hss/internal/models"
	"hss/internal/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CompanyMockGroup struct {
	mockObjects []*models.Company
}

func NewCompanyMockGroup(mockObjects ...*models.Company) CompanyMockGroup {
	return CompanyMockGroup{
		mockObjects: mockObjects,
	}
}

func (m CompanyMockGroup) loadSelf(ctx context.Context, pool *pgxpool.Pool) error {
	companyRepository := repositories.NewCompanyRepository(pool)
	for _, company := range m.mockObjects {
		err := companyRepository.InsertCompany(ctx, company)
		if err != nil {
			return fmt.Errorf("failed to load company mock: %w", err)
		}
	}
	return nil
}

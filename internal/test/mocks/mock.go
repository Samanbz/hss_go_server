package mocks

import (
	"context"
	"hss/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MockGroup interface {
	loadSelf(ctx context.Context, pool *pgxpool.Pool) error
}

func NewRelatedMocks(ctx context.Context, pool *pgxpool.Pool, companyMock *models.Company) error {
	companyMockGroup := NewCompanyMockGroup(companyMock)

	err := NewMocks(ctx, pool, companyMockGroup)
	if err != nil {
		return err
	}

	return nil
}

type Mocks struct {
	mockGroups []MockGroup
}

func NewMocks(ctx context.Context, pool *pgxpool.Pool, mockGroups ...MockGroup) error {
	mocks := &Mocks{
		mockGroups: mockGroups,
	}

	err := mocks.loadSelf(ctx, pool)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mocks) loadSelf(ctx context.Context, pool *pgxpool.Pool) error {
	for _, mockGroup := range m.mockGroups {
		err := mockGroup.loadSelf(ctx, pool)
		if err != nil {
			return err
		}
	}
	return nil
}

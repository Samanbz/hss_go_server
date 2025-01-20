package mocks

import (
	"context"
	"hss/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MockGroup interface {
	LoadSelf(ctx context.Context, pool *pgxpool.Pool) error
}

func NewRelatedMocks(ctx context.Context, pool *pgxpool.Pool,
	companyMock *models.Company, addressMock *models.Address) error {

	companyMockGroup := NewCompanyMockGroup(companyMock)
	addressMockGroup := NewAddressMockGroup(addressMock.WithForeignKey(companyMock.ID))
	//TODO Add more mock groups here

	err := NewMocks(ctx, pool, companyMockGroup, addressMockGroup)
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

	err := mocks.LoadSelf(ctx, pool)
	if err != nil {
		return err
	}

	return nil
}

func (m Mocks) LoadSelf(ctx context.Context, pool *pgxpool.Pool) error {
	for _, mockGroup := range m.mockGroups {
		err := mockGroup.LoadSelf(ctx, pool)
		if err != nil {
			return err
		}
	}
	return nil
}

// TODO: MORE SOPHISTICATED MOCKING SYSTEM WITH BASEMODEL AND REPOSITORY AS COMMON DENOMINATOR

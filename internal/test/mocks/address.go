package mocks

import (
	"context"
	"fmt"
	"hss/internal/models"
	"hss/internal/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AddressMockGroup struct {
	mockObjects []*models.Address
}

func NewAddressMockGroup(mockObjects ...*models.Address) AddressMockGroup {
	return AddressMockGroup{
		mockObjects: mockObjects,
	}
}

func (m AddressMockGroup) LoadSelf(ctx context.Context, pool *pgxpool.Pool) error {
	addressRepository := repositories.NewAddressRepository(pool)
	for _, address := range m.mockObjects {
		err := addressRepository.Create(ctx, address)
		if err != nil {
			return fmt.Errorf("failed to load address mock: %w", err)
		}
	}
	return nil
}

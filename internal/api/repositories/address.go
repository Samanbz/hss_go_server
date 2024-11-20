package repositories

import (
	"context"
	"hss/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AddressRepository struct {
	conn *pgxpool.Pool
}

func NewAddressRepository(pool *pgxpool.Pool) AddressRepository {
	return AddressRepository{conn: pool}
}

func (r *AddressRepository) InsertAddress(ctx context.Context, address *models.Address) error {
	err := address.ValidateInput()
	if err != nil {
		return err
	}

	var id int
	query := `
		INSERT INTO address (username, password, company_id, street, city, state, zip, country, latitude, longitude)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10
		) RETURNING id
	`

	err = r.conn.QueryRow(ctx, query, address.Username, address.Password, address.CompanyID, address.State, address.City, address.State, address.Zip, address.Country, address.Latitude, address.Longitude).Scan(&id)
	if err != nil {
		return err
	}

	address.ID = id

	return nil
}

package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	conn *pgxpool.Pool
}

func NewAuthRepository(pool *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{conn: pool}
}

func (r *AuthRepository) CheckCompanyByUsername(ctx context.Context, username string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM company WHERE username = $1)`
	row := r.conn.QueryRow(ctx, query, username)

	var exists bool
	err := row.Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *AuthRepository) ValidateAddressCredentials(ctx context.Context, companyUsername, addressUsername, PasswordHash string) (int, error) {
	query := `SELECT a.id AS address_id FROM company c JOIN address a ON c.id = a.company_id WHERE c.username = $1 AND a.username = $2 AND a.password = $3`
	row := r.conn.QueryRow(ctx, query, companyUsername, addressUsername, PasswordHash)

	var addressID int
	err := row.Scan(&addressID)

	if err != nil {
		return -1, err
	}

	return addressID, nil
}

func (r *AuthRepository) ValidateCompanyCredentials(ctx context.Context, companyUsername, PasswordHash string) (int, error) {
	query := `SELECT c.id AS company_id FROM company c WHERE c.username = $1 AND c.password = $2`
	row := r.conn.QueryRow(ctx, query, companyUsername, PasswordHash)

	var companyID int
	err := row.Scan(&companyID)

	if err != nil {
		return -1, err
	}

	return companyID, nil
}

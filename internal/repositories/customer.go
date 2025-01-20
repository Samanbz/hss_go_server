package repositories

import (
	"context"
	"hss/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerRepository struct {
	conn *pgxpool.Pool
}

func NewCustomerRepository(pool *pgxpool.Pool) *CustomerRepository {
	return &CustomerRepository{conn: pool}
}

func (r *CustomerRepository) Create(ctx context.Context, customer *models.Customer) error {
	err := customer.ValidateInput()
	if err != nil {
		return err
	}

	var id int
	query := `INSERT INTO customer (username, password, company_id) VALUES ($1, $2, $3) RETURNING id`

	err = r.conn.QueryRow(ctx, query, customer.Username, customer.Password, customer.CompanyID).Scan(&id)
	if err != nil {
		return err
	}

	customer.ID = id

	return nil
}

func (r *CustomerRepository) GetByID(ctx context.Context, id int) (*models.Customer, error) {
	query := `SELECT * FROM customer WHERE id = $1`
	row := r.conn.QueryRow(ctx, query, id)

	customer := models.Customer{}
	err := row.Scan(&customer.ID, &customer.Username, &customer.Password, &customer.CompanyID)
	if err != nil {
		return nil, err
	}

	err = customer.ValidateOutput()
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

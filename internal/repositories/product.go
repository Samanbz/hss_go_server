package repositories

import (
	"context"
	"hss/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	conn *pgxpool.Pool
}

func NewProductRepository(pool *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{conn: pool}
}

func (r *ProductRepository) Create(ctx context.Context, product *models.Product) error {
	err := product.ValidateInput()
	if err != nil {
		return err
	}

	var id int
	query := `INSERT INTO product (company_id, address_id, name, description, price, stock) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err = r.conn.QueryRow(ctx, query, product.CompanyID, product.AddressID, product.Name, product.Description, product.Price, product.Stock).Scan(&id)
	if err != nil {
		return err
	}

	product.ID = id

	return nil
}

func (r *ProductRepository) GetByID(ctx context.Context, id int) (*models.Product, error) {
	query := `SELECT * FROM product WHERE id = $1`
	row := r.conn.QueryRow(ctx, query, id)

	product := models.Product{}
	err := row.Scan(&product.ID, &product.CompanyID, &product.AddressID, &product.Name, &product.Description, &product.Price, &product.Stock)
	if err != nil {
		return nil, err
	}

	err = product.ValidateOutput()
	if err != nil {
		return nil, err
	}

	return &product, nil
}

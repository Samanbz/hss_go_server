package repositories

import (
	"context"
	"hss/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ServiceRepository struct {
	conn *pgxpool.Pool
}

func NewServiceRepository(pool *pgxpool.Pool) *ServiceRepository {
	return &ServiceRepository{conn: pool}
}

func (r *ServiceRepository) Create(ctx context.Context, service *models.Service) error {
	err := service.ValidateInput()
	if err != nil {
		return err
	}

	var id int
	query := `INSERT INTO service (company_id, address_id, title, price) VALUES ($1, $2, $3, $4) RETURNING id`

	err = r.conn.QueryRow(ctx, query, service.CompanyID, service.AddressID, service.Title, service.Price).Scan(&id)
	if err != nil {
		return err
	}

	service.ID = id

	return nil
}

func (r *ServiceRepository) GetByID(ctx context.Context, id int) (*models.Service, error) {
	query := `SELECT * FROM service WHERE id = $1`
	row := r.conn.QueryRow(ctx, query, id)

	service := models.Service{}
	err := row.Scan(&service.ID, &service.AddressID, &service.CompanyID, &service.Title, &service.Price)
	if err != nil {
		return nil, err
	}

	err = service.ValidateOutput()
	if err != nil {
		return nil, err
	}

	return &service, nil
}

func (r *ServiceRepository) GetAllForAddress(ctx context.Context, addressID int) (*[]models.Service, error) {
	query := `SELECT * FROM service WHERE address_id = $1`
	rows, err := r.conn.Query(ctx, query, addressID)
	if err != nil {
		return nil, err
	}

	services := []models.Service{}
	for rows.Next() {
		service := models.Service{}
		err = rows.Scan(&service.ID, &service.AddressID, &service.CompanyID, &service.Title, &service.Price)
		if err != nil {
			return nil, err
		}

		err = service.ValidateOutput()
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}

	return &services, nil
}

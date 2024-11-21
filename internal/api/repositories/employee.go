package repositories

import (
	"context"
	"hss/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepository struct {
	conn *pgxpool.Pool
}

func NewEmployeeRepository(pool *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{conn: pool}
}

func (r *EmployeeRepository) InsertEmployee(ctx context.Context, employee *models.Employee) error {
	err := employee.ValidateInput()
	if err != nil {
		return err
	}

	var id int
	query := `INSERT INTO employee (firstname, lastname, company_id, address_id, email, phone) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err = r.conn.QueryRow(ctx, query, employee.Firstname, employee.Lastname, employee.CompanyID, employee.AddressID, employee.Email, employee.Phone).Scan(&id)
	if err != nil {
		return err
	}

	employee.ID = id

	return nil
}

func (r *EmployeeRepository) GetEmployeeByID(ctx context.Context, id int) (*models.Employee, error) {
	query := `SELECT * FROM employee WHERE id = $1`
	row := r.conn.QueryRow(ctx, query, id)

	employee := models.Employee{}
	err := row.Scan(&employee.ID, &employee.Firstname, &employee.Lastname, &employee.CompanyID, &employee.AddressID, &employee.Email, &employee.Phone)
	if err != nil {
		return nil, err
	}

	err = employee.ValidateOutput()
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

package repositories

import (
	"context"
	"hss/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CompanyRepository struct {
	conn *pgxpool.Pool
}

func NewCompanyRepository(pool *pgxpool.Pool) *CompanyRepository {
	return &CompanyRepository{conn: pool}
}

func (r *CompanyRepository) Create(ctx context.Context, company *models.Company) error {
	err := company.ValidateInput()
	if err != nil {
		return err
	}

	var id int
	query := `INSERT INTO company (username, company_name, rep_firstname, rep_lastname, email, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err = r.conn.QueryRow(ctx, query, company.Username, company.CompanyName, company.RepFirstname, company.RepLastname, company.Email, company.Password).Scan(&id)
	if err != nil {
		return err
	}

	company.ID = id
	return nil
}

func (r *CompanyRepository) GetByID(ctx context.Context, id int) (*models.Company, error) {
	query := `SELECT * FROM company WHERE id = $1`
	row := r.conn.QueryRow(ctx, query, id)

	company := models.Company{}
	err := row.Scan(&company.ID, &company.Username, &company.CompanyName, &company.RepFirstname, &company.RepLastname, &company.Email, &company.OTPSecret, &company.Password)
	if err != nil {
		return nil, err
	}
	company.ValidateOutput()

	return &company, nil
}

func (r *CompanyRepository) UpdateByID(ctx context.Context, id int, company *models.Company) error {
	err := company.ValidateInput()
	if err != nil {
		return err
	}

	query := `UPDATE company SET username = $1, company_name = $2, rep_firstname = $3, rep_lastname = $4, email = $5, password = $6 WHERE id = $7`
	_, err = r.conn.Exec(ctx, query, company.Username, company.CompanyName, company.RepFirstname, company.RepLastname, company.Email, company.Password, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *CompanyRepository) DeleteByID(ctx context.Context, id int) error {
	query := `DELETE FROM company WHERE id = $1`
	_, err := r.conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *CompanyRepository) GetByUsername(ctx context.Context, username string) (*models.Company, error) {
	query := `SELECT * FROM company WHERE username = $1`
	row := r.conn.QueryRow(ctx, query, username)

	company := models.Company{}
	err := row.Scan(&company.ID, &company.Username, &company.CompanyName, &company.RepFirstname, &company.RepLastname, &company.Email, &company.OTPSecret, &company.Password)
	if err != nil {
		return nil, err
	}
	company.ValidateOutput()

	return &company, nil
}

func (r *CompanyRepository) GetAll(ctx context.Context) ([]models.Company, error) {
	query := `SELECT * FROM company`
	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	companies := []models.Company{}
	for rows.Next() {
		company := models.Company{}
		err := rows.Scan(&company.ID, &company.Username, &company.CompanyName, &company.RepFirstname, &company.RepLastname, &company.Email, &company.OTPSecret, &company.Password)
		if err != nil {
			return nil, err
		}
		company.ValidateOutput()
		companies = append(companies, company)
	}
	return companies, nil
}

package repositories

import (
	"context"
	"hss/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CompanyRepository struct {
	conn *pgxpool.Pool
}

func NewCompanyRepository(pool *pgxpool.Pool) CompanyRepository {
	return CompanyRepository{conn: pool}
}

func (r CompanyRepository) InsertCompany(ctx context.Context, company *models.Company) error {
	company.ValidateInput()
	query := `INSERT INTO companies (username, company_name, rep_firstname, rep_lastname, email, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var id uint32
	err := r.conn.QueryRow(ctx, query, company.Username, company.CompanyName, company.RepFirstname, company.RepLastname, company.Email, company.Password).Scan(&id)

	company.ID = &id

	if err != nil {
		return err
	}
	return nil
}

func (r CompanyRepository) GetCompanyByID(ctx context.Context, id uint32) (*models.Company, error) {
	query := `SELECT * FROM companies WHERE id = $1`
	row := r.conn.QueryRow(ctx, query, id)

	company := models.Company{}
	err := row.Scan(&company.ID, &company.Username, &company.CompanyName, &company.RepFirstname, &company.RepLastname, &company.Email, &company.Password)
	company.ValidateOutput()
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (r CompanyRepository) GetCompanyByUsername(ctx context.Context, username string) (*models.Company, error) {
	query := `SELECT * FROM companies WHERE username = $1`
	row := r.conn.QueryRow(ctx, query, username)

	company := models.Company{}
	err := row.Scan(&company.ID, &company.Username, &company.CompanyName, &company.RepFirstname, &company.RepLastname, &company.Email, &company.Password)
	company.ValidateOutput()
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (r CompanyRepository) GetAllCompanies(ctx context.Context) ([]models.Company, error) {
	query := `SELECT * FROM companies`
	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	companies := []models.Company{}
	for rows.Next() {
		company := models.Company{}
		err := rows.Scan(&company.ID, &company.Username, &company.CompanyName, &company.RepFirstname, &company.RepLastname, &company.Email, &company.Password)
		company.ValidateOutput()
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}
	return companies, nil
}

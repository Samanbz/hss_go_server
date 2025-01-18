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

func (r *CompanyRepository) InsertCompany(ctx context.Context, company *models.Company) error {
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

func (r *CompanyRepository) GetCompanyByID(ctx context.Context, id int) (*models.Company, error) {
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

func (r *CompanyRepository) GetCompanyByUsername(ctx context.Context, username string) (*models.Company, error) {
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

func (r *CompanyRepository) GetAllCompanies(ctx context.Context) ([]models.Company, error) {
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

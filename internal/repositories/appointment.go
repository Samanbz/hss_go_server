package repositories

import (
	"context"
	"hss/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AppointmentRepository struct {
	conn *pgxpool.Pool
}

func NewAppointmentRepository(pool *pgxpool.Pool) *AppointmentRepository {
	return &AppointmentRepository{conn: pool}
}

func (r *AppointmentRepository) Create(ctx context.Context, appointment *models.Appointment) error {
	err := appointment.ValidateInput()
	if err != nil {
		return err
	}

	var id int
	query := `INSERT INTO appointment (start, end, company_id, address_id, employee_id, service_id, customer_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	err = r.conn.QueryRow(ctx, query, appointment.Start, appointment.End, appointment.CompanyID, appointment.AddressID, appointment.EmployeeID, appointment.ServiceID, appointment.CustomerID).Scan(&id)
	if err != nil {
		return err
	}

	appointment.ID = id

	return nil
}

func (r *AppointmentRepository) GetByID(ctx context.Context, id int) (*models.Appointment, error) {
	query := `SELECT * FROM appointment WHERE id = $1`
	row := r.conn.QueryRow(ctx, query, id)

	appointment := models.Appointment{}
	err := row.Scan(&appointment.ID, &appointment.Start, &appointment.End, &appointment.CompanyID, &appointment.AddressID, &appointment.EmployeeID, &appointment.ServiceID, &appointment.CustomerID)
	appointment.ValidateOutput()
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (r *AppointmentRepository) GetAllForCompany(ctx context.Context, company_id int) (*[]models.Appointment, error) {
	query := `SELECT * FROM appointment WHERE company_id = $1`
	rows, err := r.conn.Query(ctx, query, company_id)
	if err != nil {
		return nil, err
	}

	appointments := []models.Appointment{}
	for rows.Next() {
		appointment := models.Appointment{}
		err = rows.Scan(&appointment.ID, &appointment.Start, &appointment.End, &appointment.CompanyID, &appointment.AddressID, &appointment.EmployeeID, &appointment.ServiceID, &appointment.CustomerID)
		if err != nil {
			return nil, err
		}
		appointment.ValidateOutput()
		appointments = append(appointments, appointment)
	}

	return &appointments, nil
}

package models

import (
	"time"
)

type Appointment struct {
	ID         *int       `json:"id"`
	Start      *time.Time `json:"start" validate:"required"`
	End        *time.Time `json:"end" validate:"required"`
	CompanyID  *int       `json:"user_id" validate:"required"`
	AddressID  *int       `json:"address_id" validate:"required"`
	EmployeeID *int       `json:"employee_id"`
	ServiceID  *int       `json:"service_id"`
	CustomerID *int       `json:"customer_id" validate:"required"`
}

func NewAppointment(
	start, end *time.Time,
	companyID, addressID, employeeID, serviceID, customerID *int) *Appointment {

	return &Appointment{
		Start:      start,
		End:        end,
		CompanyID:  companyID,
		AddressID:  addressID,
		EmployeeID: employeeID,
		ServiceID:  serviceID,
		CustomerID: customerID,
	}
}

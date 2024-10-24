package models

import (
	"time"
)

type Appointment struct {
	ID         *uint32    `json:"id"`
	Start      *time.Time `json:"start" validate:"required"`
	End        *time.Time `json:"end" validate:"required"`
	CompanyID  *uint32    `json:"user_id" validate:"required"`
	AddressID  *uint32    `json:"address_id" validate:"required"`
	EmployeeID *uint32    `json:"employee_id"`
	ServiceID  *uint32    `json:"service_id"`
	CustomerID *uint32    `json:"customer_id" validate:"required"`
}

func NewAppointment(
	start, end *time.Time,
	companyID, addressID, employeeID, serviceID, customerID *uint32) *Appointment {

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

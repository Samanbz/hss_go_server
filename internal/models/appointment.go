package models

import (
	"encoding/json"
	"hss/pkg/validation"
	"time"
)

type Appointment struct {
	ID         int       `json:"id"`
	Start      time.Time `json:"start" validate:"required"`
	End        time.Time `json:"end" validate:"required"`
	CompanyID  int       `json:"user_id" validate:"required"`
	AddressID  int       `json:"address_id" validate:"required"`
	EmployeeID int       `json:"employee_id"`
	ServiceID  int       `json:"service_id"`
	CustomerID int       `json:"customer_id" validate:"required"`
}

func NewAppointment(
	start, end time.Time,
	companyID, addressID, employeeID, serviceID, customerID int) *Appointment {

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

func NewAppointmentFromJSON(data []byte) (*Appointment, error) {
	appointment := Appointment{}
	err := json.Unmarshal(data, &appointment)
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (a *Appointment) ValidateInput() error {
	return validation.Validate.Struct(a)
}

func (a *Appointment) ValidateOutput() {
	validation.Validate.Struct(a)
}

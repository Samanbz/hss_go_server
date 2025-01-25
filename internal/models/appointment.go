package models

import (
	"encoding/json"
	"hss/internal/utils"
	"hss/internal/utils/validation"
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

func (a Appointment) ToJSON() []byte {
	jsonData, _ := json.Marshal(a)
	return jsonData
}

func (a *Appointment) FromJSON(jsonData []byte) error {
	return json.Unmarshal(jsonData, a)
}

func (a *Appointment) ValidateInput() error {
	return validation.GetValidator().Struct(a)
}

func (a *Appointment) ValidateOutput() error {
	return validation.GetValidator().Struct(a)
}

func (a Appointment) Hash() string {
	return utils.Hash(string(a.ToJSON()))
}

func (a Appointment) Equals(other Appointment) bool {
	return a.Hash() == other.Hash()
}

func (a Appointment) WithForeignKeys(companyID, addressID, employeeID, serviceID, customerID int) *Appointment {
	a.CompanyID = companyID
	a.AddressID = addressID
	a.EmployeeID = employeeID
	a.ServiceID = serviceID
	a.CustomerID = customerID

	return &a
}

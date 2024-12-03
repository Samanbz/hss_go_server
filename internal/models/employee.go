package models

import (
	"encoding/json"
	"hss/internal/utils/validation"
)

type Employee struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	AddressID int    `json:"address_id" validate:"required"`
	CompanyID int    `json:"company_id" validate:"required"`
	Email     string `json:"email" validate:"email"`
	Phone     string `json:"phone" validate:"e164"`
}

func NewEmployee(firstname, lastname, email, phone string, addressID, companyID int) *Employee {

	return &Employee{
		Firstname: firstname,
		Lastname:  lastname,
		AddressID: addressID,
		CompanyID: companyID,
		Email:     email,
		Phone:     phone,
	}
}

func NewEmployeeFromJSON(jsonData []byte) (*Employee, error) {
	var employee Employee
	err := json.Unmarshal(jsonData, &employee)

	return &employee, err
}

func (e *Employee) ValidateInput() error {
	return validation.Validate.StructExcept(e, "ID")
}

func (e *Employee) ValidateOutput() error {
	return validation.Validate.Struct(e)
}

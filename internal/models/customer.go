package models

import (
	"encoding/json"
	"hss/pkg/utils/security"
	"hss/pkg/validation"
)

type Customer struct {
	ID        int    `json:"id"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	CompanyID string `json:"user_id" validate:"required"`
}

func NewCustomer(username, password, companyID string) *Customer {

	hashedPassword := security.HashPassword(password)

	return &Customer{
		Username:  username,
		Password:  hashedPassword,
		CompanyID: companyID,
	}
}

func NewCustomerFromJSON(jsonData []byte) (*Customer, error) {
	var customer Customer
	err := json.Unmarshal(jsonData, &customer)

	return &customer, err
}

func (c *Customer) ValidateInput() error {
	return validation.Validate.StructExcept(c, "ID")
}

func (c *Customer) ValidateOutput() error {
	return validation.Validate.Struct(c)
}

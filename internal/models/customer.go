package models

import (
	"hss/pkg/utils"
)

type Customer struct {
	ID        *uint32 `json:"id"`
	Username  *string `json:"username" validate:"required"`
	Password  *string `json:"password" validate:"required"`
	CompanyID *string `json:"user_id" validate:"required"`
}

func NewCustomer(username, password, companyID string) Customer {

	hashedPassword := utils.HashPassword(&password)

	return Customer{
		Username:  &username,
		Password:  &hashedPassword,
		CompanyID: &companyID,
	}
}

package models

import "hss/internal/models/base"

type Employee struct {
	ID            *uint32               `json:"id"`
	Firstname     *string               `json:"firstname" validate:"required"`
	Lastname      *string               `json:"lastname" validate:"required"`
	AddressID     *string               `json:"address_id" validate:"required"`
	Email         *string               `json:"email" validate:"email"`
	Phone         *string               `json:"phone" validate:"phone"`
	WorkingHourss *[7]base.WorkingHours `json:"working_hours"`
}

func NewEmployee(firstname, lastname, addressID, email, phone string, workingHours *[7]*base.WorkingHours) Employee {

	return Employee{
		Firstname: &firstname,
		Lastname:  &lastname,
		AddressID: &addressID,
		Email:     &email,
		Phone:     &phone,
	}
}

package models

type Employee struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	AddressID string `json:"address_id" validate:"required"`
	Email     string `json:"email" validate:"email"`
	Phone     string `json:"phone" validate:"phone"`
}

func NewEmployee(firstname, lastname, addressID, email, phone string) Employee {

	return Employee{
		Firstname: firstname,
		Lastname:  lastname,
		AddressID: addressID,
		Email:     email,
		Phone:     phone,
	}
}

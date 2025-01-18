package models

import (
	"encoding/json"
	"hss/internal/utils/validation"
)

type Address struct {
	ID        int     `json:"id"`
	Username  string  `json:"username" validate:"required"`
	Password  string  `json:"password" validate:"required,sha256"`
	CompanyID int     `json:"company_id" validate:"required"`
	Street    string  `json:"street" validate:"required"`
	City      string  `json:"city" validate:"required"`
	State     string  `json:"state" validate:"required"`
	Zip       string  `json:"zip" validate:"required"`
	Country   string  `json:"country" validate:"required"`
	Latitude  float32 `json:"latitude" validate:"latitude"`
	Longitude float32 `json:"longitude" validate:"longitude"`
}

func NewAddress(
	username, password, street, city, state, zip, country string,
	latitude, longitude float32,
	companyID int) *Address {

	return &Address{
		Username:  username,
		Password:  password,
		CompanyID: companyID,
		Street:    street,
		City:      city,
		State:     state,
		Zip:       zip,
		Country:   country,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

func NewAddressFromJSON(jsonData []byte) (*Address, error) {
	var address Address
	err := json.Unmarshal(jsonData, &address)

	return &address, err
}

func (u Address) ToJSON() []byte {
	jsonData, _ := json.Marshal(u)
	return jsonData
}

func (u *Address) FromJSON(jsonData []byte) error {
	return json.Unmarshal(jsonData, u)
}

func (u *Address) ValidateInput() error {
	return validation.Validate.StructExcept(u, "ID")
}

func (u *Address) ValidateOutput() error {
	return validation.Validate.Struct(u)
}

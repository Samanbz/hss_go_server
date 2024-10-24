package models

import "hss/internal/models/base"

type Address struct {
	ID            *uint32               `json:"id"`
	Username      *string               `json:"username" validate:"required"`
	Password      *string               `json:"password" validate:"required,sha256"`
	CompanyID     *uint32               `json:"company_id" validate:"required"`
	Street        *string               `json:"street" validate:"required"`
	City          *string               `json:"city" validate:"required"`
	State         *string               `json:"state" validate:"required"`
	Zip           *string               `json:"zip" validate:"required"`
	Country       *string               `json:"country" validate:"required"`
	Latitude      *float32              `json:"latitude" validate:"latitude"`
	Longitude     *float32              `json:"longitude" validate:"longitude"`
	WorkingHourss *[7]base.WorkingHours `json:"working_hours"`
}

func NewAddress(
	username, password, street, city, state, zip, country *string,
	latitude, longitude *float32,
	companyID *uint32) *Address {

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

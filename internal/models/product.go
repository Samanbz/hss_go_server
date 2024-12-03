package models

import (
	"encoding/json"
	"hss/internal/utils/validation"
)

type Product struct {
	ID          int     `json:"id"`
	CompanyID   string  `json:"company_id" validate:"required"`
	AddressID   string  `json:"address_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
}

func NewProduct(
	companyID, addressID, name, description string,
	price float32, stock int) *Product {

	return &Product{
		CompanyID:   companyID,
		AddressID:   addressID,
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}
}

func NewProductFromJSON(jsonData []byte) (*Product, error) {
	var product Product
	err := json.Unmarshal(jsonData, &product)

	return &product, err
}

func (p *Product) ValidateInput() error {
	return validation.Validate.StructExcept(p, "ID")
}

func (p *Product) ValidateOutput() error {
	return validation.Validate.Struct(p)
}

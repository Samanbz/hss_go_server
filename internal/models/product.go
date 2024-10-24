package models

type Product struct {
	ID          *uint32  `json:"id"`
	CompanyID   *string  `json:"company_id" validate:"required"`
	AddressID   *string  `json:"address_id" validate:"required"`
	Name        *string  `json:"name" validate:"required"`
	Description *string  `json:"description" validate:"required"`
	Price       *float32 `json:"price" validate:"required"`
	Stock       *int     `json:"stock" validate:"required"`
}

func NewProduct(
	companyID, addressID, name, description *string,
	price *float32, stock *int) *Product {

	return &Product{
		CompanyID:   companyID,
		AddressID:   addressID,
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}
}

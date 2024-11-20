package models

type Service struct {
	ID        *int     `json:"id"`
	AddressID *string  `json:"address_id" validate:"required"`
	CompanyID *string  `json:"company_id" validate:"required"`
	Title     *string  `json:"title" validate:"required"`
	Price     *float32 `json:"price" validate:"required"`
}

func NewService(addressID, companyID, title *string, price *float32) *Service {
	return &Service{
		AddressID: addressID,
		CompanyID: companyID,
		Title:     title,
		Price:     price,
	}
}

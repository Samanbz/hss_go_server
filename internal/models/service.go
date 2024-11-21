package models

import (
	"encoding/json"
	"hss/pkg/validation"
)

type Service struct {
	ID        int     `json:"id"`
	AddressID string  `json:"address_id" validate:"required"`
	CompanyID string  `json:"company_id" validate:"required"`
	Title     string  `json:"title" validate:"required"`
	Price     float32 `json:"price" validate:"required"`
}

func NewService(addressID, companyID, title string, price float32) *Service {
	return &Service{
		AddressID: addressID,
		CompanyID: companyID,
		Title:     title,
		Price:     price,
	}
}

func NewServiceFromJSON(jsonData []byte) (*Service, error) {
	var service Service
	err := json.Unmarshal(jsonData, &service)

	return &service, err
}

func (s *Service) ValidateInput() error {
	return validation.Validate.StructExcept(s, "ID")
}

func (s *Service) ValidateOutput() error {
	return validation.Validate.Struct(s)
}

package models

import (
	"encoding/json"
	"hss/internal/utils"
	"hss/internal/utils/validation"
)

type Service struct {
	ID        int     `json:"id"`
	AddressID int     `json:"address_id" validate:"required"`
	CompanyID int     `json:"company_id" validate:"required"`
	Title     string  `json:"title" validate:"required"`
	Price     float32 `json:"price" validate:"required"`
}

func (s Service) ToJSON() []byte {
	jsonData, _ := json.Marshal(s)
	return jsonData
}

func (s *Service) FromJSON(jsonData []byte) error {
	return json.Unmarshal(jsonData, s)
}

func (s *Service) ValidateInput() error {
	return validation.GetValidator().StructExcept(s, "ID")
}

func (s *Service) ValidateOutput() error {
	return validation.GetValidator().Struct(s)
}

func (s Service) Hash() string {
	return utils.Hash(string(s.ToJSON()))
}

func (s Service) Equals(other Service) bool {
	return s.Hash() == other.Hash()
}

func (s Service) WithForeignKeys(companyID, addressID int) *Service {
	s.CompanyID = companyID
	s.AddressID = addressID

	return &s
}

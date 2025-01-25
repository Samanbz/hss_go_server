package models

import (
	"encoding/json"
	"hss/internal/utils"
	"hss/internal/utils/validation"
)

type Product struct {
	ID          int     `json:"id"`
	CompanyID   int     `json:"company_id" validate:"required"`
	AddressID   int     `json:"address_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
}

func (p Product) ToJSON() []byte {
	jsonData, _ := json.Marshal(p)
	return jsonData
}

func (p *Product) FromJSON(jsonData []byte) error {
	return json.Unmarshal(jsonData, p)
}

func (p *Product) ValidateInput() error {
	return validation.GetValidator().StructExcept(p, "ID")
}

func (p *Product) ValidateOutput() error {
	return validation.GetValidator().Struct(p)
}

func (p Product) Hash() string {
	return utils.Hash(string(p.ToJSON()))
}

func (p Product) Equals(other Product) bool {
	return p.Hash() == other.Hash()
}

func (p Product) WithForeignKeys(companyID, addressID int) *Product {
	p.CompanyID = companyID
	p.AddressID = addressID

	return &p
}

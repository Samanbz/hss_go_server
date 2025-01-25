package models

import (
	"encoding/json"
	"hss/internal/utils"
	"hss/internal/utils/validation"
)

type Employee struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	AddressID int    `json:"address_id" validate:"required"`
	CompanyID int    `json:"company_id" validate:"required"`
	Email     string `json:"email" validate:"email"`
	Phone     string `json:"phone" validate:"e164"`
}

func (e Employee) ToJSON() []byte {
	jsonData, _ := json.Marshal(e)
	return jsonData
}

func (e *Employee) FromJSON(jsonData []byte) error {
	return json.Unmarshal(jsonData, e)
}

func (e *Employee) ValidateInput() error {
	return validation.GetValidator().StructExcept(e, "ID")
}

func (e *Employee) ValidateOutput() error {
	return validation.GetValidator().Struct(e)
}

func (e Employee) Hash() string {
	return utils.Hash(string(e.ToJSON()))
}

func (e Employee) Equals(other Employee) bool {
	return e.Hash() == other.Hash()
}

func (e Employee) WithForeignKeys(companyID, addressID int) *Employee {
	e.CompanyID = companyID
	e.AddressID = addressID
	return &e
}

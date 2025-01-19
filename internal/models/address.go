package models

import (
	"encoding/json"
	"hss/internal/utils"
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

func (u Address) ToJSON() []byte {
	jsonData, _ := json.Marshal(u)
	return jsonData
}

func (u *Address) FromJSON(jsonData []byte) error {
	return json.Unmarshal(jsonData, u)
}

func (u *Address) ValidateInput() error {
	return validation.GetValidator().StructExcept(u, "ID")
}

func (u *Address) ValidateOutput() error {
	return validation.GetValidator().Struct(u)
}

func (u Address) Hash() string {
	return utils.Hash(string(u.ToJSON()))
}

func (u Address) Equals(other Address) bool {
	return u.Hash() == other.Hash()
}

package models

import (
	"encoding/json"
	"hss/internal/utils"
	"hss/internal/utils/validation"
)

type Company struct {
	ID           int     `json:"id" validate:"required"`
	Username     string  `json:"username" validate:"required,gt=5"`
	CompanyName  string  `json:"company_name" validate:"required"`
	RepFirstname string  `json:"rep_firstname" validate:"required"`
	RepLastname  string  `json:"rep_lastname" validate:"required"`
	Email        string  `json:"email" validate:"required,email"`
	OTPSecret    *string `json:"otp_secret"`
	Password     string  `json:"password" validate:"required,sha256"`
}

func (u Company) ToJSON() []byte {
	jsonData, _ := json.Marshal(u)
	return jsonData
}

func (u *Company) FromJSON(jsonData []byte) error {
	return json.Unmarshal(jsonData, u)
}

func (u Company) ValidateInput() error {
	return validation.GetValidator().StructExcept(u, "ID")
}

func (u Company) ValidateOutput() error {
	return validation.GetValidator().Struct(u)
}

func (u Company) Hash() string {
	return utils.Hash(string(u.ToJSON()))
}

func (u Company) Equals(c Comparable) bool {
	return u.Hash() == c.Hash()
}

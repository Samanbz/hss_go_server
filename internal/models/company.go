package models

import (
	"encoding/json"
	"hss/pkg/utils/security"
	"hss/pkg/validation"
)

type Company struct {
	ID           int    `json:"id" validate:"required"`
	Username     string `json:"username" validate:"required,gt=5"`
	CompanyName  string `json:"company_name" validate:"required"`
	RepFirstname string `json:"rep_firstname" validate:"required"`
	RepLastname  string `json:"rep_lastname" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	OTPSecret    string `json:"otp_secret"`
	Password     string `json:"password" validate:"required,sha256"`
}

func NewCompany(username, companyName, repFirstname, repLastname, email, password string) Company {

	hashedPassword := security.HashPassword(password)

	return Company{
		Username:     username,
		CompanyName:  companyName,
		RepFirstname: repFirstname,
		RepLastname:  repLastname,
		Email:        email,
		OTPSecret:    "",
		Password:     hashedPassword,
	}
}

func NewCompanyFromJSON(jsonData []byte) (*Company, error) {
	var company Company
	err := json.Unmarshal(jsonData, &company)
	company.Password = security.HashPassword(company.Password)

	return &company, err
}

func (u *Company) ValidateInput() error {
	return validation.Validate.StructExcept(u, "ID")
}

func (u *Company) ValidateOutput() error {
	return validation.Validate.Struct(u)
}

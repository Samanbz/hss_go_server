package models

import (
	"fmt"
	"hss/pkg/utils"
	"hss/pkg/validation"
)

type Company struct {
	ID           *uint32 `json:"id" validate:"required"`
	Username     *string `json:"username" validate:"required,gt=5"`
	CompanyName  *string `json:"company_name" validate:"required"`
	RepFirstname *string `json:"rep_firstname" validate:"required"`
	RepLastname  *string `json:"rep_lastname" validate:"required"`
	Email        *string `json:"email" validate:"required,email"`
	OTPSecret    *string `json:"otp_secret"`
	Password     *string `json:"password" validate:"required,sha256"`
}

func NewCompany(username, companyName, repFirstname, repLastname, email, password string) Company {

	hashedPassword := utils.HashPassword(&password)

	return Company{
		Username:     &username,
		CompanyName:  &companyName,
		RepFirstname: &repFirstname,
		RepLastname:  &repLastname,
		Email:        &email,
		OTPSecret:    nil,
		Password:     &hashedPassword,
	}
}

func (u *Company) ValidateInput() error {
	return validation.Validate.StructExcept(u, "ID")
}

func (u *Company) ValidateOutput() error {
	return validation.Validate.Struct(u)
}

func (u Company) ToString() (string, error) {
	if err := u.ValidateOutput(); err != nil {
		return "NOT VALID", err
	}

	id := *u.ID
	username := *u.Username
	companyName := *u.CompanyName
	repFirstname := *u.RepFirstname
	repLastname := *u.RepLastname
	email := *u.Email
	otpSecret := "NULL"
	if u.OTPSecret != nil {
		otpSecret = *u.OTPSecret
	}
	password := *u.Password

	return fmt.Sprintf(
		"ID: %d\nUsername: %s\nCompanyName: %s\nRepFirstname: %s\nRepLastname: %s\nEmail: %s\nOTPSecret: %s\nPassword: %s",
		id, username, companyName, repFirstname, repLastname, email, otpSecret, password), nil
}

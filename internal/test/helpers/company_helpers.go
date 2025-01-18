package helpers

import (
	"errors"
	"fmt"
	"hss/internal/models"
	"strings"
)

func CheckCompanyID(companyID int, expectedID int) error {
	if companyID != expectedID {
		return fmt.Errorf("expected company ID %d, got %d", expectedID, companyID)
	}
	return nil
}

func CheckCompanyUsername(companyUsername string, expectedUsername string) error {
	if companyUsername != expectedUsername {
		return fmt.Errorf("expected company username %s, got %s", expectedUsername, companyUsername)
	}
	return nil
}

func CheckCompanyCompanyName(companyName string, expectedName string) error {
	if companyName != expectedName {
		return fmt.Errorf("expected company name %s, got %s", expectedName, companyName)
	}
	return nil
}

func CheckCompanyRepFirstname(repFirstname string, expectedRepFirstname string) error {
	if repFirstname != expectedRepFirstname {
		return fmt.Errorf("expected company representative first name %s, got %s", expectedRepFirstname, repFirstname)
	}
	return nil
}

func CheckCompanyRepLastname(repLastname string, expectedRepLastname string) error {
	if repLastname != expectedRepLastname {
		return fmt.Errorf("expected company representative last name %s, got %s", expectedRepLastname, repLastname)
	}
	return nil
}

func CheckCompanyEmail(email string, expectedEmail string) error {
	if email != expectedEmail {
		return fmt.Errorf("expected company email %s, got %s", expectedEmail, email)
	}
	return nil
}

func CheckCompanyPassword(password string, expectedPassword string) error {
	if password != expectedPassword {
		return fmt.Errorf("expected company password %s, got %s", expectedPassword, password)
	}
	return nil
}

func CheckCompany(company *models.Company, expected *models.Company, checkID bool) error {
	var errorMessages []string

	if company.ID == 0 {
		errorMessages = append(errorMessages, "company ID is 0")
	}
	if err := CheckCompanyID(company.ID, expected.ID); checkID && err != nil {
		errorMessages = append(errorMessages, err.Error())
	}
	if err := CheckCompanyUsername(company.Username, expected.Username); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}
	if err := CheckCompanyCompanyName(company.CompanyName, expected.CompanyName); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}
	if err := CheckCompanyRepFirstname(company.RepFirstname, expected.RepFirstname); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}
	if err := CheckCompanyRepLastname(company.RepLastname, expected.RepLastname); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}
	if err := CheckCompanyEmail(company.Email, expected.Email); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}
	if err := CheckCompanyPassword(company.Password, expected.Password); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if len(errorMessages) > 0 {
		return errors.New(strings.Join(errorMessages, "\n"))
	}
	return nil
}

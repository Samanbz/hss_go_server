package mock_data

import "hss/internal/models"

var (
	MockEmployee = models.Employee{
		Firstname: "Alice",
		Lastname:  "Johnson",
		AddressID: 1,
		CompanyID: 101,
		Email:     "alice.johnson@example.com",
		Phone:     "+12345678901",
	}

	MockEmployee2 = models.Employee{
		Firstname: "Bob",
		Lastname:  "Smith",
		AddressID: 2,
		CompanyID: 102,
		Email:     "bob.smith@example.com",
		Phone:     "+19876543210",
	}

	MockEmployee3 = models.Employee{
		Firstname: "Charlie",
		Lastname:  "Brown",
		AddressID: 3,
		CompanyID: 103,
		Email:     "charlie.brown@example.com",
		Phone:     "+11234567890",
	}

	MockEmployee4 = models.Employee{
		Firstname: "Diana",
		Lastname:  "Prince",
		AddressID: 4,
		CompanyID: 104,
		Email:     "diana.prince@example.com",
		Phone:     "+10987654321",
	}
)

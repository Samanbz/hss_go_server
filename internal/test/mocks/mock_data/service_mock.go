package mock_data

import "hss/internal/models"

var (
	MockService1 = models.Service{
		AddressID: 1,
		CompanyID: 101,
		Title:     "Haircut and Styling",
		Price:     45.00,
	}

	MockService2 = models.Service{
		AddressID: 2,
		CompanyID: 102,
		Title:     "Manicure and Pedicure",
		Price:     60.00,
	}

	MockService3 = models.Service{
		AddressID: 3,
		CompanyID: 103,
		Title:     "Facial Treatment",
		Price:     75.00,
	}
)

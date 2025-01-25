package mock_data

import "hss/internal/models"

var (
	MockProduct1 = models.Product{
		ID:          1,
		CompanyID:   101,
		AddressID:   1001,
		Name:        "Shampoo - Revitalizing Blend",
		Description: "For dry scalp, reduces flakes, adds moisture.",
		Price:       15.99,
		Stock:       50,
	}

	MockProduct2 = models.Product{
		ID:          2,
		CompanyID:   102,
		AddressID:   1002,
		Name:        "Nourishing Face Cream",
		Description: "For dry skin, hydrates deeply, adds glow.",
		Price:       25.50,
		Stock:       30,
	}

	MockProduct3 = models.Product{
		ID:          3,
		CompanyID:   103,
		AddressID:   1003,
		Name:        "Manicure Kit - Deluxe Edition",
		Description: "All-in-one, for salon-quality nails at home.",
		Price:       35.75,
		Stock:       20,
	}
)

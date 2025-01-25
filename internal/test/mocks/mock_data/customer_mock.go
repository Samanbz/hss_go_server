package mock_data

import "hss/internal/models"

var (
	MockCustomer1 = models.Customer{
		Username:  "johndoe",
		Password:  "4af478de7863e056ed5cef911b86abd92c95023b08335052974fb037a0b94101",
		CompanyID: 1,
	}

	MockCustomer2 = models.Customer{
		Username:  "janedoe",
		Password:  "b4cf69166575c2048f8c73f57238079c47a69ce5c4e1e3579c1bfbf7e4f78d58",
		CompanyID: 2,
	}

	MockCustomer3 = models.Customer{
		Username:  "mikesmith",
		Password:  "97767a213f255a4bb99505cf2bb69153174899805a6fec01a9b6aa4c9e97a849",
		CompanyID: 3,
	}
)

package mock_data

import (
	"hss/internal/models"
	"hss/internal/utils"
)

var (
	MockAppointment1 = models.Appointment{
		Start:      utils.InNDaysAt(1, 10, 0),
		End:        utils.InNDaysAt(1, 11, 0),
		CompanyID:  1,
		AddressID:  1,
		EmployeeID: 1,
		ServiceID:  1,
		CustomerID: 1,
	}

	MockAppointment2 = models.Appointment{
		Start:      utils.InNDaysAt(2, 10, 0),
		End:        utils.InNDaysAt(2, 11, 0),
		CompanyID:  2,
		AddressID:  2,
		EmployeeID: 2,
		ServiceID:  2,
		CustomerID: 2,
	}

	MockAppointment3 = models.Appointment{
		Start:      utils.InNDaysAt(3, 10, 0),
		End:        utils.InNDaysAt(3, 11, 0),
		CompanyID:  3,
		AddressID:  3,
		EmployeeID: 3,
		ServiceID:  3,
		CustomerID: 3,
	}
)

package repositories

type Repositories struct {
	CompanyRepository     *CompanyRepository
	AddressRepository     *AddressRepository
	EmployeeRepository    *EmployeeRepository
	AppointmentRepository *AppointmentRepository
	ServiceRepository     *ServiceRepository
	CustomerRepository    *CustomerRepository
	ProductRepository     *ProductRepository
	AuthRepository        *AuthRepository
}

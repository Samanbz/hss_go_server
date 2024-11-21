package routes

import (
	"hss/internal/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func CompanyRoutes(app *fiber.App, CompanyHandler *handlers.CompanyHandler) {
	app.Post("/company", CompanyHandler.InsertCompany)
	app.Get("/company", CompanyHandler.GetAllCompanies)
}

func AddressRoutes(app *fiber.App, AddressHandler *handlers.AddressHandler) {
	app.Post("/address", AddressHandler.InsertAddress)
}

func EmployeeRoutes(app *fiber.App, EmployeeHandler *handlers.EmployeeHandler) {
	app.Post("/employee", EmployeeHandler.InsertEmployee)
	app.Get("/employee/:id", EmployeeHandler.GetEmployeeByID)
}

func AppointmentRoutes(app *fiber.App, AppointmentHandler *handlers.AppointmentHandler) {
	app.Post("/appointment", AppointmentHandler.InsertAppointment)
	app.Get("/appointment/:id", AppointmentHandler.GetAppointmentByID)
	app.Get("/appointment/company/:id", AppointmentHandler.GetAppointmentsByCompanyID)
}

func ServiceRoutes(app *fiber.App, ServiceHandler *handlers.ServiceHandler) {
	app.Post("/service", ServiceHandler.InsertService)
	app.Get("/service/:id", ServiceHandler.GetServiceByID)
	app.Get("/service/address/:id", ServiceHandler.GetServicesByAddressID)
}

func CustomerRoutes(app *fiber.App, CustomerHandler *handlers.CustomerHandler) {
	app.Post("/customer", CustomerHandler.InsertCustomer)
	app.Get("/customer/:id", CustomerHandler.GetCustomerByID)
}

func ProductRoutes(app *fiber.App, ProductHandler *handlers.ProductHandler) {
	app.Post("/product", ProductHandler.InsertProduct)
	app.Get("/product/:id", ProductHandler.GetProductByID)
}

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

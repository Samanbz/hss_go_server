package main

import (
	"hss/internal/api/routes"
	"hss/internal/database"
	"hss/pkg/utils/singleton"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	pool := database.InitDB()
	defer pool.Close()

	CompanyHandler, AddressHandler, EmployeeHandler, AppointmentHandler, err := singleton.InitSingletons(pool)
	if err != nil {
		panic(err)
	}

	routes.CompanyRoutes(app, CompanyHandler)
	routes.AddressRoutes(app, AddressHandler)
	routes.EmployeeRoutes(app, EmployeeHandler)
	routes.AppointmentRoutes(app, AppointmentHandler)

	app.Get("/", handleRoot)
	app.Listen(":3000")
}

func handleRoot(c *fiber.Ctx) error {
	return c.SendString("Hello, World!!")
}

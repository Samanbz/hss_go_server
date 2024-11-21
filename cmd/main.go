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

	RequestHandlers, err := singleton.InitSingletons(pool)
	if err != nil {
		panic(err)
	}

	routes.CompanyRoutes(app, RequestHandlers.CompanyHandler)
	routes.AddressRoutes(app, RequestHandlers.AddressHandler)
	routes.EmployeeRoutes(app, RequestHandlers.EmployeeHandler)
	routes.AppointmentRoutes(app, RequestHandlers.AppointmentHandler)
	routes.ServiceRoutes(app, RequestHandlers.ServiceHandler)
	routes.CustomerRoutes(app, RequestHandlers.CustomerHandler)
	routes.ProductRoutes(app, RequestHandlers.ProductHandler)

	app.Get("/", handleRoot)
	app.Listen(":3000")
}

func handleRoot(c *fiber.Ctx) error {
	return c.SendString("Hello, World!!")
}

package main

import (
	"hss/internal/database"
	"hss/internal/routes"
	"hss/internal/utils/singleton"

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

	routes.InitRoutes(app, RequestHandlers)

	app.Listen(":3000")
}

func handleRoot(c *fiber.Ctx) error {
	return c.SendString("Hello, World!!")
}

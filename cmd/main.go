package main

import (
	"context"
	"fmt"
	"hss/internal/handlers"
	"hss/internal/models"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	app := fiber.New()

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_DOMAIN"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	//TODO move to new location
	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		log.Fatalf("Error parsing config: %v\n", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Error creating pool: %v\n", err)
	}
	defer pool.Close()
	//

	app.Post("/company", handlers.InsertCompany)
	company := models.NewCompany(
		"username", "companyName", "repFirstname", "repLastname", "some.email@gmail.com", "password")

	str, _ := company.ToString()
	fmt.Printf("%s\n", str)

	app.Get("/", handleRoot)
	app.Listen(":3000")
}

func handleRoot(c *fiber.Ctx) error {
	return c.SendString("Hello, World!!")
}

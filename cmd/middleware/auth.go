package middleware

import (
	"hss/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx, authHandler handlers.AuthHandler) error {
	valid, err := authHandler.ValidateToken(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "cannot validate token",
		})
	}

	if !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	return c.Next()
}

package handlers

import (
	"context"
	"fmt"
	"hss/internal/models"
	"hss/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) ValidateCredentials(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	authRequest, err := models.NewAuthenticationRequestFromJSON(c.Body())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot parse JSON, %v", err),
		})
	}

	authResponse, err := h.authService.ValidateCredentials(ctx, authRequest)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot validate credentials, %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(authResponse)
}

func (h *AuthHandler) ValidateToken(c *fiber.Ctx) (bool, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	token := c.Get("Authorization")
	if token == "" {
		return false, nil
	}

	authReq := models.AuthorizationRequest{Token: token}
	_, err := h.authService.ValidateToken(ctx, &authReq)
	if err != nil {
		return false, err
	}

	return true, nil
}

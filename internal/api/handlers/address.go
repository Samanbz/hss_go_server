package handlers

import (
	"context"
	"fmt"
	"hss/internal/api/services"
	"hss/internal/models"

	"github.com/gofiber/fiber/v2"
)

type AddressHandler struct {
	addressService services.AddressService
}

func NewAddressHandler(addressService services.AddressService) AddressHandler {
	return AddressHandler{addressService: addressService}
}

func (h *AddressHandler) InsertAddress(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	address, err := models.NewAddressFromJSON(c.Body())

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot parse JSON, %v", err),
		})
	}

	if err := h.addressService.InsertAddress(ctx, address); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot insert address, %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(address)
}

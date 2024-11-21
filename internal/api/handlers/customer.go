package handlers

import (
	"context"
	"fmt"
	"hss/internal/api/services"
	"hss/internal/models"

	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	customerService *services.CustomerService
}

func NewCustomerHandler(customerService *services.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}

func (h *CustomerHandler) InsertCustomer(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	customer, err := models.NewCustomerFromJSON(c.Body())

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot parse JSON, %v", err),
		})
	}

	if err := h.customerService.InsertCustomer(ctx, customer); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot insert customer, %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

func (h *CustomerHandler) GetCustomerByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot parse id, %v", err),
		})
	}

	customer, err := h.customerService.GetCustomerByID(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot get customer, %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

package handlers

import (
	"context"
	"fmt"
	"hss/internal/models"
	"hss/internal/services"

	"github.com/gofiber/fiber/v2"
)

type ServiceHandler struct {
	serviceService *services.ServiceService // I'm sorry you had to read that
}

func NewServiceHandler(serviceService *services.ServiceService) *ServiceHandler {
	return &ServiceHandler{serviceService: serviceService}
}

func (h *ServiceHandler) InsertService(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	service, err := models.NewServiceFromJSON(c.Body())

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot parse JSON, %v", err),
		})
	}

	if err := h.serviceService.InsertService(ctx, service); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot insert service, %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(service)
}

func (h *ServiceHandler) GetServiceByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot parse id, %v", err),
		})
	}

	service, err := h.serviceService.GetServiceByID(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot get service, %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(service)
}

func (h *ServiceHandler) GetServicesByAddressID(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	addressID, err := c.ParamsInt("addressID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot parse addressID, %v", err),
		})
	}

	services, err := h.serviceService.GetServicesByAddressID(ctx, addressID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot get services, %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(services)
}

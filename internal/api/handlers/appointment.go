package handlers

import (
	"context"
	"fmt"
	"hss/internal/api/services"
	"hss/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AppointmentHandler struct {
	appointmentService *services.AppointmentService
}

func NewAppointmentHandler(appointmentService *services.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{appointmentService: appointmentService}
}

func (h *AppointmentHandler) InsertAppointment(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	appointment, err := models.NewAppointmentFromJSON(c.Body())

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot parse JSON, %v", err),
		})
	}

	if err := h.appointmentService.InsertAppointment(ctx, appointment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot insert appointment, %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(appointment)
}

func (h *AppointmentHandler) GetAppointmentByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("invalid appointment ID, %v", err),
		})
	}

	appointment, err := h.appointmentService.GetAppointmentByID(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot get appointment, %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(appointment)
}

func (h *AppointmentHandler) GetAppointmentsByCompanyID(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("invalid appointment ID, %v", err),
		})
	}

	appointments, err := h.appointmentService.GetAppointmentsByCompanyID(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot get appointments, %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(appointments)
}

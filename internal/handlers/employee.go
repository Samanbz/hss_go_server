package handlers

import (
	"context"
	"fmt"
	"hss/internal/models"
	"hss/internal/services"

	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct {
	employeeService *services.EmployeeService
}

func NewEmployeeHandler(employeeService *services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{employeeService: employeeService}
}

func (h *EmployeeHandler) InsertEmployee(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	employee, err := models.NewEmployeeFromJSON(c.Body())

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot parse JSON, %v", err),
		})
	}

	if err := h.employeeService.InsertEmployee(ctx, employee); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot insert employee, %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(employee)
}

func (h *EmployeeHandler) GetEmployeeByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot parse id, %v", err),
		})
	}

	employee, err := h.employeeService.GetEmployeeByID(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot get employee, %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(employee)
}

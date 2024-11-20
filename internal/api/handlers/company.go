package handlers

import (
	"context"
	"fmt"
	"hss/internal/api/services"
	"hss/internal/models"

	"github.com/gofiber/fiber/v2"
)

type CompanyHandler struct {
	companyService services.CompanyService
}

func NewCompanyHandler(companyService services.CompanyService) CompanyHandler {
	return CompanyHandler{companyService: companyService}
}

func (h *CompanyHandler) InsertCompany(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	company, err := models.NewCompanyFromJSON(c.Body())

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot parse JSON, %v", err),
		})
	}

	if err := h.companyService.InsertCompany(ctx, company); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot insert company, %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(company)
}

func (h *CompanyHandler) GetAllCompanies(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	companies, err := h.companyService.GetAllCompanies(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("cannot get companies, %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(companies)
}

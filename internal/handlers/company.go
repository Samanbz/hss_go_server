package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func InsertCompany(fctx *fiber.Ctx) error {
	//ctx, cancel := context.WithCancel(context.Background())
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	//TODO extract json data
	return nil
}

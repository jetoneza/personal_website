package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jetoneza/personal_website/pkg/application"
)

type Handler struct {
	App *application.Application
}

func NewHandler(app *application.Application) *Handler {
	return &Handler{
		App: app,
	}
}

func (h *Handler) HealthCheck(c *fiber.Ctx) error {
	// TODO: Add status codes to constants
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "jetrooper.me API",
	})
}

package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handler) Signup(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}
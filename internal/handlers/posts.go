package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jetoneza/personal_website/internal/models"
)

func (h *Handler) GetAllPosts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset := (page - 1) * limit

	var posts []models.Post

	results := h.App.DB.Limit(limit).Offset(offset).Find(&posts)

	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": results.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"results": len(posts),
		"posts":   posts,
	})
}

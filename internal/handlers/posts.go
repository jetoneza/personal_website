package handlers

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jetoneza/personal_website/internal/models"
	"github.com/jetoneza/personal_website/internal/schema"
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
		"data":    posts,
	})
}

func (h *Handler) CreatePost(c *fiber.Ctx) error {
	payload := new(schema.CreatePostSchema)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// TODO: Validate data

	post := models.Post{
		Title:    payload.Title,
		Content:  payload.Content,
		Category: payload.Category,
	}

	result := h.App.DB.Create(&post)

	if result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate entry") {
		// TODO: Add constants for messages
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Title already exist, please use another title"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   post,
	})
}

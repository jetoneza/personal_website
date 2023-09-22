package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jetoneza/personal_website/internal/models"
	"github.com/jetoneza/personal_website/internal/schema"
	"github.com/jetoneza/personal_website/pkg/utils"
)

func (h *Handler) GetPost(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var post models.Post

	result := h.App.DB.First(&post, "id = ?", id)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("%v", result.Error),
		})
	}

	// Convert md to html
	post.ConvertContentToHtml()

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   post,
	})
}

func (h *Handler) GetAllPosts(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	offset := (page - 1) * limit

	var posts []models.Post

	results := h.App.DB.Limit(limit).Offset(offset).Find(&posts)

	if results.Error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status":  "error",
			"message": results.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"results": len(posts),
		"data":    posts,
	})
}

func (h *Handler) CreatePost(ctx *fiber.Ctx) error {
	body := new(schema.CreatePostSchema)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "error": err.Message})
	}

	post := models.Post{
		Title:    body.Title,
		Content:  body.Content,
		Category: body.Category,
	}

	result := h.App.DB.Create(&post)

	if result.Error != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "error": result.Error.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   post,
	})
}

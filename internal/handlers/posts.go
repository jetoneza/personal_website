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

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   post,
	})
}

func (h *Handler) GetPostBySlug(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")

	var post models.Post

	result := h.App.DB.First(&post, "slug = ?", slug)

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

	results := h.App.DB.Order("created_at desc").Limit(limit).Offset(offset).Find(&posts)

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

func (h *Handler) GetAllPublishedPosts(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	offset := (page - 1) * limit

	var posts []models.Post

	results := h.App.DB.Order("created_at desc").Limit(limit).Offset(offset).Find(&posts, "published = true")

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
		Title:           body.Title,
		Slug:            body.Slug,
		Description:     body.Description,
		Content:         body.Content,
		Category:        body.Category,
		MetaTitle:       body.MetaTitle,
		MetaDescription: body.MetaDescription,
		MetaKeywords:    body.MetaKeywords,
		MetaImageUrl:    body.MetaImageUrl,
		Published:       body.Published,
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

func (h *Handler) UpdatePost(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var post models.Post

	result := h.App.DB.First(&post, "id = ?", id)

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("%v", result.Error),
		})
	}

	body := new(schema.CreatePostSchema)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "error": err.Message})
	}

	post.Title = body.Title
	post.Slug = body.Slug
	post.Description = body.Description
	post.Content = body.Content
	post.Category = body.Category
	post.Published = body.Published
	post.MetaTitle = body.MetaTitle
	post.MetaDescription = body.MetaDescription
	post.MetaKeywords = body.MetaKeywords
	post.MetaImageUrl = body.MetaImageUrl

	saveResult := h.App.DB.Save(&post)

	if saveResult.Error != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "error": saveResult.Error.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   post,
	})
}

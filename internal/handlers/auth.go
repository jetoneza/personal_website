package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jetoneza/personal_website/internal/models"
	"github.com/jetoneza/personal_website/internal/schema"
	"github.com/jetoneza/personal_website/pkg/utils"
	"gorm.io/gorm"
)

const AUTH_INVALID_CREDS_ERROR = "Invalid email or password"

func (h *Handler) Signup(ctx *fiber.Ctx) error {
	body := new(schema.SignupSchema)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "error": err.Message})
	}

	err := h.App.DB.Model(&models.User{}).Take(&struct{ ID string }{}, "email = ?", body.Email).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status": "fail",
			"error":  "Email already exists",
		})
	}

	user := &models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: utils.GenerateHashedPassword(body.Password),
	}

	result := h.App.DB.Create(&user)

	if result.Error != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "error": result.Error.Error()})
	}

	token := utils.GenerateAuthToken(&utils.AuthTokenPayload{
		ID:    user.ID,
		Email: user.Email,
	})

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": &schema.AuthResponse{
			User: &schema.UserResponse{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			},
			Token: token,
		},
	})
}

func (h *Handler) Login(ctx *fiber.Ctx) error {
	body := new(schema.LoginSchema)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "error": err.Message})
	}

	var user models.User

	result := h.App.DB.First(&user, "email = ?", body.Email)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "fail",
			"error":  AUTH_INVALID_CREDS_ERROR,
		})
	}

	if err := utils.VerifyPassword(body.Password, user.Password); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "fail",
			"error":  AUTH_INVALID_CREDS_ERROR,
		})
	}

	token := utils.GenerateAuthToken(&utils.AuthTokenPayload{
		ID:    user.ID,
		Email: user.Email,
	})

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": &schema.AuthResponse{
			User: &schema.UserResponse{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			},
			Token: token,
		},
	})
}

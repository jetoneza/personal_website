package handlers

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jetoneza/personal_website/internal/models"
	"github.com/jetoneza/personal_website/internal/schema"
	"github.com/jetoneza/personal_website/pkg/utils"
	"gorm.io/gorm"
)

const AUTH_INVALID_CREDS_ERROR = "Invalid email or password"

func (h *Handler) RegisterUser(ctx *fiber.Ctx) error {
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

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func (h *Handler) LoginUser(ctx *fiber.Ctx) error {
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

	token, expiry := utils.GenerateAuthToken(&utils.AuthTokenPayload{
		ID:    user.ID,
		Email: user.Email,
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		Expires:  expiry,
		Secure:   true,
		HTTPOnly: true,
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

func (h *Handler) LogoutUser(ctx *fiber.Ctx) error {
	ctx.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "session-ended",
		Path:     "/",
		Expires:  time.Now().Add(time.Second * 10),
		Secure:   true,
		HTTPOnly: true,
	})

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Logged out successfully",
	})
}

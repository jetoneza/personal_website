package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jetoneza/personal_website/internal/models"
	// "github.com/jetoneza/personal_website/internal/schema"
	// "github.com/jetoneza/personal_website/pkg/utils"
)

func (h *Handler) GetUserSettings(ctx *fiber.Ctx) error {
	jwtData := ctx.Locals("user").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)

	var settings models.UserSetting

	result := h.App.DB.First(&settings, "user_id = ?", claims["id"])

	if result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("%v", result.Error),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   settings,
	})
}

func (h *Handler) UpsertUserSettings(ctx *fiber.Ctx) error {
	// jwtData := ctx.Locals("user").(*jwt.Token)
	// claims := jwtData.Claims.(jwt.MapClaims)
	//
	// body := new(schema.CreateUserSettingSchema)
	//
	// if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "error": err.Message})
	// }

	// TODO: Implement upsert logic here

	return nil
}

package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (h *Handlers) GetUserSettings(ctx *fiber.Ctx) error {
	jwtData := ctx.Locals("user").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)

	var settings UserSettings

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

funct (h *Handlers) UpsertUserSettings(ctx *fiber.Ctx) error {
  jwtData := ctx.Locals("user").(*jwt.Token)
  claims := jwtData.Claims.(jwt.MapClaims)

  body := new (schema.UserSettings)

  if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
    return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "error": err.Message})
  }
})

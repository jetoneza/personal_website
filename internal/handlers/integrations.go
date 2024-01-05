package handlers

import (
	"encoding/json"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/jetoneza/personal_website/internal/models"
	"github.com/jetoneza/personal_website/pkg/config"
	"gorm.io/gorm"
)

type Error struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Code     string `json:"code"`
}

type StravaError struct {
	Message string  `json:"message"`
	Errors  []Error `json:"errors"`
}

func (h *Handler) ProcessStravaExchangeToken(ctx *fiber.Ctx) error {
	code := ctx.Query("code")

	if code == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"error":  "Invalid code",
		})
	}

	queryParams := url.Values{
		"client_id":     {config.STRAVA_CLIENT_ID},
		"client_secret": {config.STRAVA_CLIENT_SECRET},
		"code":          {code},
		"grant_type":    {"authorization_code"},
	}

	// Get token from strava
	url := "https://www.strava.com/oauth/token?" + queryParams.Encode()
	agent := fiber.Post(url)
	statusCode, body, errors := agent.Bytes()

	if len(errors) > 0 || statusCode != 200 {
		var error StravaError
		json.Unmarshal(body, &error)

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "fail",
			"statusCode": statusCode,
			"message":    error.Message,
			"errors":     error.Errors,
		})
	}

	// Save integration tokens
	var integration models.Integration

	var result *gorm.DB

	result = h.App.DB.First(&integration, "name = ?", "Strava")
	// Create if not exists
	if result.Error != nil {
		json.Unmarshal(body, &integration)

		integration.Name = "Strava"

		result = h.App.DB.Create(&integration)
	} else {
		// Update if exists
		json.Unmarshal(body, &integration)

		result = h.App.DB.Save(&integration)
	}

	if result.Error != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "error": result.Error.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   integration,
	})
}

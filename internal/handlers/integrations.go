package handlers

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/jetoneza/personal_website/internal/models"
	"github.com/jetoneza/personal_website/internal/schema/strava"
	"github.com/jetoneza/personal_website/pkg/config"
	"github.com/jetoneza/personal_website/pkg/utils"
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

func (h *Handler) CreateStravaActivity(ctx *fiber.Ctx) error {
	body := new(strava.CreateActivity)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "error": err.Message})
	}

	token, err := h.GetStravaAccessToken()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "error": err.Error()})
	}

	// Create activity in strava
	url := "https://www.strava.com/api/v3/activities"
	agent := fiber.Post(url)
	agent.Set("Authorization", "Bearer "+token)

	args := fiber.AcquireArgs()
	args.Set("name", body.Name)
	args.Set("type", body.Type)
	args.Set("sport_type", body.SportType)
	args.Set("trainer", body.Trainer)
	args.Set("distance", body.Distance)
	args.Set("elapsed_time", body.ElapsedTime)
	args.Set("start_date_local", body.StartDateLocal)

	agent.Form(args)

	statusCode, responseBody, errors := agent.Bytes()

	fiber.ReleaseArgs(args)

	if len(errors) > 0 {
		var error StravaError
		json.Unmarshal(responseBody, &error)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "fail",
			"statusCode": statusCode,
			"message":    error.Message,
			"errors":     error.Errors,
		})
	}

	type Response struct {
		Name string `json:"name"`
		Id   int `json:"id"`
		// TODO: Add more response fields
	}
	var response Response

	json.Unmarshal(responseBody, &response)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   response,
	})
}

func (h *Handler) GetStravaAccessToken() (string, error) {
	var integration models.Integration

	result := h.App.DB.First(&integration, "name = ?", "Strava")

	if result.Error != nil {
		return "", result.Error
	}

	if !integration.IsExpired() {
		return integration.AccessToken, nil
	}

	queryParams := url.Values{
		"client_id":     {config.STRAVA_CLIENT_ID},
		"client_secret": {config.STRAVA_CLIENT_SECRET},
		"refresh_token": {integration.RefreshToken},
		"grant_type":    {"refresh_token"},
	}

	url := "https://www.strava.com/oauth/token?" + queryParams.Encode()
	agent := fiber.Post(url)

	statusCode, body, responseErrors := agent.Bytes()
	if len(responseErrors) > 0 || statusCode != 200 {
		var error StravaError
		json.Unmarshal(body, &error)

		return "", errors.New(error.Message)
	}

	var response models.Integration
	json.Unmarshal(body, &response)

	integration.TokenType = response.TokenType
	integration.ExpiresAt = response.ExpiresAt
	integration.ExpiresIn = response.ExpiresIn
	integration.RefreshToken = response.RefreshToken
	integration.AccessToken = response.AccessToken

	saveResult := h.App.DB.Save(&integration)

	if saveResult.Error != nil {
		return "", saveResult.Error
	}

	return integration.AccessToken, nil
}

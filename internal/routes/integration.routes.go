package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jetoneza/personal_website/internal/handlers"
	"github.com/jetoneza/personal_website/internal/middlewares"
)

func IntegrationRoutes(router fiber.Router, handlers *handlers.Handler) {
	r := router.Group("/integrations")

	r.Get("/strava", middlewares.BasicAuth(), handlers.ProcessStravaExchangeToken)
	r.Post("/strava/activities", middlewares.BasicAuth(), handlers.CreateStravaActivity)
}

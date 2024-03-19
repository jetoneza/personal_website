package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jetoneza/personal_website/internal/handlers"
	"github.com/jetoneza/personal_website/internal/middlewares"
)

func UserRoutes(router fiber.Router, handlers *handlers.Handler) {
	r := router.Group("/users")

	r.Get("/user_settings", middlewares.AuthMiddleware(), handlers.GetUserSettings)
}

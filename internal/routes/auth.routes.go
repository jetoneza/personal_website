package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jetoneza/personal_website/internal/handlers"
)

func AuthRoutes(router fiber.Router, handlers *handlers.Handler) {
	r := router.Group("/auth")

	r.Post("/register", handlers.RegisterUser)
	r.Post("/login", handlers.LoginUser)
	r.Post("/logout", handlers.LogoutUser)
}

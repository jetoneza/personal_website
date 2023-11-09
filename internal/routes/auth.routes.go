package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jetoneza/personal_website/internal/handlers"
	"github.com/jetoneza/personal_website/internal/middlewares"
)

func AuthRoutes(router fiber.Router, handlers *handlers.Handler) {
	r := router.Group("/auth")

	r.Post("/register", middlewares.BasicAuth(), handlers.RegisterUser)
	r.Post("/login", handlers.LoginUser)
	r.Post("/logout", handlers.LogoutUser)
	r.Get("/session", middlewares.AuthMiddleware(), handlers.GetSessionInfo)
}

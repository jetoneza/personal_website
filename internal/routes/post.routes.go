package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jetoneza/personal_website/internal/handlers"
)

func PostRoutes(router fiber.Router, handlers *handlers.Handler) {
	r := router.Group("/posts")

	r.Get("/", handlers.GetAllPosts)
	r.Post("/", handlers.CreatePost)
	r.Get("/:id", handlers.GetPost)
}
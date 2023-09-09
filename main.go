package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jetoneza/personal_website/internal/handlers"
	"github.com/jetoneza/personal_website/pkg/application"
	"github.com/jetoneza/personal_website/web"
)

// TODO: Move to configs using .env
const (
	database   = "posts.sqlite"
	appName    = "Jet Ordaneza Personal Website"
	port       = ":3000"
	apiVersion = "v1"
)

func main() {
	app := application.New()
	app.InitializeDB(database)

	fiberApp := fiber.New(fiber.Config{
		AppName: appName,
	})
	defer fiberApp.Shutdown()

	// Middlewares
	fiberApp.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Routes
	handlers := handlers.NewHandler(app)

	v1 := fiberApp.Group("/api/" + apiVersion)
	v1.Get("/healthcheck", handlers.HealthCheck)

	postsRoute := v1.Group("/posts")
	postsRoute.Get("/", handlers.GetAllPosts)
	postsRoute.Post("/", handlers.CreatePost)

	// Serve static files
	fiberApp.All("/*", filesystem.New(filesystem.Config{
		Root:         web.Build(),
		NotFoundFile: "index.html",
		Index:        "index.html",
	}))

	// Go!
	if err := fiberApp.Listen(port); err != nil {
		panic(err)
	}
}

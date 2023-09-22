package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/jetoneza/personal_website/internal/handlers"
	"github.com/jetoneza/personal_website/internal/middlewares"
	"github.com/jetoneza/personal_website/internal/routes"
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
  middlewares.SetupMiddlewares(fiberApp)

	// Handlers
	handlers := handlers.NewHandler(app)

	v1 := fiberApp.Group("/api/" + apiVersion)
	v1.Get("/healthcheck", handlers.HealthCheck)

	// TODO: Add auth middleware to POST endpoints
  routes.PostRoutes(v1, handlers)

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

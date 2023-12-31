package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/jetoneza/personal_website/internal/handlers"
	"github.com/jetoneza/personal_website/internal/middlewares"
	"github.com/jetoneza/personal_website/internal/routes"
	"github.com/jetoneza/personal_website/pkg/application"
	"github.com/jetoneza/personal_website/pkg/config"
	"github.com/jetoneza/personal_website/web"
)

const (
	apiVersion = "v1"
)

func main() {
	app := application.New()
	app.InitializeDB(config.DB)

	fiberApp := fiber.New(fiber.Config{
		AppName: config.APP_NAME,
	})
	defer fiberApp.Shutdown()

	// Middlewares
	middlewares.SetupMiddlewares(fiberApp)

	// Handlers
	handlers := handlers.NewHandler(app)

	v1 := fiberApp.Group("/api/" + apiVersion)
	v1.Get("/healthcheck", handlers.HealthCheck)

	routes.AuthRoutes(v1, handlers)
	routes.PostRoutes(v1, handlers)
	routes.EventRoutes(v1, handlers)
	routes.IntegrationRoutes(v1, handlers)

	// Serve static files
	fiberApp.All("/*", filesystem.New(filesystem.Config{
		Root:         web.Build(),
		NotFoundFile: "index.html",
		Index:        "index.html",
	}))

	// Go!
	if err := fiberApp.Listen(fmt.Sprintf(":%v", config.PORT)); err != nil {
		panic(err)
	}
}

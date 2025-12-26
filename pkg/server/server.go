package server

import (
	_ "project/docs" // Import generated docs
	"project/pkg/config"
	"project/pkg/errorhandler"
	"project/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

type Server struct {
	App *fiber.App
}

func New(cfg *config.Config) *Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorhandler.FiberErrorHandler,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(middleware.Trace())
	app.Use(middleware.RequestLogger())

	// Cors Config
	corsConfig := cors.Config{
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH,TRACE",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		AllowCredentials: true,
	}
	if cfg.CorsAllowedOrigins == "*" {
		corsConfig.AllowOriginsFunc = func(origin string) bool {
			return true
		}
	} else {
		corsConfig.AllowOrigins = cfg.CorsAllowedOrigins
	}

	app.Use(cors.New(corsConfig))

	// Swagger UI
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Generic Routes
	app.Get("/", hello)
	app.Get("/health", healthCheck)

	return &Server{App: app}
}

func (s *Server) Run(addr string) error {
	return s.App.Listen(addr)
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World! Try /api/todos")
}

func healthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
		"app":    "fiber-demo",
	})
}

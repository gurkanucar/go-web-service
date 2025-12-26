package bootstrap

import (
	"fmt"
	"log/slog"
	"project/internal/todo"
	"project/internal/todo/usecase"
	"project/pkg/config"
	"project/pkg/database"
	"project/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Container struct {
	// Extras
	Validator *validator.Validator
	DB        *gorm.DB

	// Repositories
	TodoRepo todo.Repository

	// UseCases
	TodoService todo.UseCase

	// Handlers
	TodoHandler *todo.Handler
}

func NewContainer() *Container {
	c := &Container{}

	c.initExtras()
	c.initDatabase()
	c.initRepositories()
	c.initUseCases()
	c.initHandlers()

	return c
}

func (c *Container) initExtras() {
	c.Validator = validator.New()
}

func (c *Container) initDatabase() {
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	db, err := database.NewConnection(cfg)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	c.DB = db

	if err := database.AutoMigrate(c.DB); err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}

	if err := database.Seed(c.DB); err != nil {
		slog.Warn("Failed to seed database", "error", err)
	}
}

func (c *Container) initRepositories() {
	c.TodoRepo = todo.NewRepository(c.DB)
}

func (c *Container) initUseCases() {
	c.TodoService = usecase.NewUseCase(c.TodoRepo)
}

func (c *Container) initHandlers() {
	c.TodoHandler = todo.NewHandler(c.TodoService, c.Validator)
}

func (c *Container) RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")
	todo.RegisterRoutes(api, c.TodoHandler)
}

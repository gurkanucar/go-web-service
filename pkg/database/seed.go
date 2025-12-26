package database

import (
	"log/slog"
	"project/internal/todo"

	"gorm.io/gorm"
)

// Seed populates the database with initial data if empty
func Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&todo.Todo{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		slog.Info("Database already seeded", "count", count)
		return nil
	}

	slog.Info("Seeding database with sample data...")

	todos := []todo.Todo{
		{Title: "Learn Go Fiber", Completed: true},
		{Title: "Implement GORM", Completed: true},
		{Title: "Add Pagination", Completed: true},
		{Title: "Deploy Application", Completed: false},
		{Title: "Set up Docker containers", Completed: false},
		{Title: "Configure Redis cache", Completed: true},
		{Title: "Write unit tests", Completed: false},
		{Title: "Add authentication middleware", Completed: true},
		{Title: "Implement JWT tokens", Completed: true},
		{Title: "Create API documentation", Completed: false},
		{Title: "Set up CI/CD pipeline", Completed: false},
		{Title: "Add logging system", Completed: true},
		{Title: "Implement rate limiting", Completed: false},
		{Title: "Configure CORS", Completed: true},
		{Title: "Add input validation", Completed: true},
		{Title: "Set up PostgreSQL", Completed: true},
		{Title: "Create database migrations", Completed: true},
		{Title: "Add error handling", Completed: true},
		{Title: "Implement file upload", Completed: false},
		{Title: "Add email notifications", Completed: false},
		{Title: "Set up monitoring", Completed: false},
		{Title: "Configure environment variables", Completed: true},
		{Title: "Add health check endpoint", Completed: true},
		{Title: "Implement graceful shutdown", Completed: false},
		{Title: "Add request timeout", Completed: true},
	}

	if err := db.Create(&todos).Error; err != nil {
		slog.Error("Failed to seed database", "error", err)
		return err
	}

	slog.Info("Database seeded successfully")
	return nil
}

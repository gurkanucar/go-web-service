package database

import (
	"log/slog"
	"project/internal/todo"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	slog.Info("Running database migrations...")

	models := []interface{}{
		&todo.Todo{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		slog.Error("Failed to run migrations", "error", err)
		return err
	}

	slog.Info("Database migrations completed successfully")
	return nil
}

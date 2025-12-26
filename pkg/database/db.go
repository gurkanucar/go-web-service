package database

import (
	"fmt"
	"log/slog"
	"strings"
	"time"

	"project/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConnection(cfg *config.Config) (*gorm.DB, error) {
	driver := cfg.DBDriver
	if driver == "" {
		dsnLower := strings.ToLower(cfg.DBDSN)
		if strings.Contains(dsnLower, "postgres") {
			driver = "postgres"
		} else if strings.Contains(dsnLower, "mysql") {
			driver = "mysql"
		} else {
			driver = "sqlite"
		}
	}

	slog.Info("Connecting to database...", "driver", driver)

	var db *gorm.DB
	var err error

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	switch driver {
	case "postgres":
		db, err = gorm.Open(postgres.Open(cfg.DBDSN), gormConfig)
	case "mysql":
		db, err = gorm.Open(mysql.Open(cfg.DBDSN), gormConfig)
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(cfg.DBDSN), gormConfig)
	default:
		return nil, fmt.Errorf("unsupported DB_DRIVER: %s", driver)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)

	slog.Info("Database connected successfully")
	return db, nil
}

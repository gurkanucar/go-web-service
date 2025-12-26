package main

import (
	"log/slog"
	"os"
	"project/pkg/bootstrap"
	"project/pkg/config"
	"project/pkg/logger"
	"project/pkg/server"
)

// @title           Project API
// @version         1.0
// @description     An example API service

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter "Bearer" followed by a space and your JWT token. Example: Bearer eyJhbGc...

func main() {
	// Load Configuration
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		os.Exit(1)
	}

	// Initialize Logger
	log := logger.New(cfg.AppEnv == "production")
	slog.SetDefault(log)

	// Initialize Dependency Container
	container := bootstrap.NewContainer()

	// Initialize Server
	srv := server.New(cfg)

	// Register Routes
	container.RegisterRoutes(srv.App)

	// Run Server
	addr := cfg.ServerAddr()
	log.Info("Server starting", "addr", addr)
	if err := srv.Run(addr); err != nil {
		log.Error("Server failed", "error", err)
	}
}

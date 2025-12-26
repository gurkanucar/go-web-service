package config

import (
	"fmt"
	"log/slog"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type SMTPConfig struct {
	Host     string `env:"SMTP_HOST"`
	Port     string `env:"SMTP_PORT"`
	User     string `env:"SMTP_USER"`
	Password string `env:"SMTP_PASSWORD"`
	From     string `env:"SMTP_FROM"`
}

type Config struct {
	ServerHost string `env:"SERVER_HOST"`
	ServerPort string `env:"SERVER_PORT"`
	AppEnv     string `env:"APP_ENV" envDefault:"development"`

	DBDriver string `env:"DB_DRIVER"`
	DBDSN    string `env:"DB_DSN"`

	CorsAllowedOrigins string `env:"CORS_ALLOWED_ORIGINS" envDefault:"*"`

	SMTP SMTPConfig
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		slog.Debug("No .env file found, relying on system environment variables")
	}

	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &cfg, nil
}

func (c *Config) ServerAddr() string {
	return fmt.Sprintf("%s:%s", c.ServerHost, c.ServerPort)
}

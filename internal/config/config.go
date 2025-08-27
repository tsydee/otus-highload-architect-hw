package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DB       DB
	Security Security
	HTTP     HTTPConfig
	Logger   Logger
}

type DB struct {
	URI string `env:"DB_URI"`
}

type Security struct {
	SecretKey string `env:"SECRET_KEY"`
}

type HTTPConfig struct {
	Port int `env:"HTTP_PORT"`
}

type Logger struct {
	LogLevel string `env:"LOG_LEVEL"`
}

func Parse() (*Config, error) {
	if os.Getenv("ENV") != "PRODUCTION" && os.Getenv("ENV") != "STAGING" {
		_ = godotenv.Load()
	}
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("load config from env: %w", err)
	}

	return &cfg, nil
}

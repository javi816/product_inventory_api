package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	AppEnv      string
	AppPort     string
	DatabaseURL string
}

func Load() (*Config, error) {

	_ = godotenv.Load()

	config := &Config{
		AppName:     os.Getenv("APP_NAME"),
		AppEnv:      os.Getenv("APP_ENV"),
		AppPort:     os.Getenv("APP_PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	if config.AppName == "" {
		return nil, errors.New("APP_NAME is required")
	}

	if config.AppEnv == "" {
		return nil, errors.New("APP_ENV is required")
	}

	if config.AppPort == "" {
		return nil, errors.New("APP_PORT is required")
	}

	if config.DatabaseURL == "" {
		return nil, errors.New("DATABASE_URL is required")
	}

	return config, nil
}

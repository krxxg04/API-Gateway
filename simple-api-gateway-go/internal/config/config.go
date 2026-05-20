package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                string
	JWTSecret           string
	AuthServiceURL      string
	UserServiceURL      string
	InventoryServiceURL string
	PaymentServiceURL   string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		Port:                getEnv("GATEWAY_PORT", "8080"),
		JWTSecret:           os.Getenv("JWT_SECRET"),
		AuthServiceURL:      getEnv("AUTH_SERVICE_URL", "http://localhost:8081"),
		UserServiceURL:      getEnv("USER_SERVICE_URL", "http://localhost:8082"),
		InventoryServiceURL: getEnv("INVENTORY_SERVICE_URL", "http://localhost:8083"),
		PaymentServiceURL:   getEnv("PAYMENT_SERVICE_URL", "http://localhost:8084"),
	}

	if cfg.JWTSecret == "" {
		return nil, errors.New("JWT_SECRET is required")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

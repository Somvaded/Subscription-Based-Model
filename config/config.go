package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
}
// LoadConfig loads configuration from environment variables or .env file
func LoadConfig()(*Config , error){
	err := godotenv.Load(".env")
	if err != nil {
		return nil,fmt.Errorf("error loading .env file: %v", err)
	}
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/subscription_service"),
		Port:        getEnv("PORT", "8080"),
	}, nil
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
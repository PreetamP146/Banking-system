package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// LoadConfig loads configuration from environment variables and returns a Config struct.
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return &Config{
		AppPort:    getENV("APP_PORT"),
		DBHost:     getENV("DB_HOST"),
		DBPort:     getENV("DB_PORT"),
		DBUser:     getENV("DB_USER"),
		DBPassword: getENV("DB_PASSWORD"),
		DBName:     getENV("DB_NAME"),
	}
}

// getENV retrieves the value of the environment variable named by the key.
func getENV(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}

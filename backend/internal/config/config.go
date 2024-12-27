package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config holds all configuration for our application
type Config struct {
	GeminiAPIKey string
	Port         string
	// Database configuration for future use
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

// New returns a new Config struct
func New() *Config {
	// Load .env file
	if err := loadEnv(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	return &Config{
		GeminiAPIKey: getEnvVar("GEMINI_API_KEY", ""),
		Port:         getEnvVar("PORT", "8080"),
		DBHost:       getEnvVar("DB_HOST", "localhost"),
		DBPort:       getEnvVar("DB_PORT", "5432"),
		DBName:       getEnvVar("DB_NAME", "trivia_db"),
		DBUser:       getEnvVar("DB_USER", "postgres"),
		DBPassword:   getEnvVar("DB_PASSWORD", "postgres"),
	}
}

// loadEnv attempts to load the .env file from the current directory
// and parent directories until it finds one or reaches the root
func loadEnv() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return godotenv.Load(envPath)
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return nil
}

// Helper function to read an environment variable or return a default value
func getEnvVar(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv string
	Port string
	DatabaseURL string
	JWTSecret string
}

func LoadConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// Load environment-specific .env file
	envFile := ".env." + env
	if _, err := os.Stat(envFile); err == nil {
		log.Println("📦 Loading environment file:", envFile)
		if err := godotenv.Load(envFile); err != nil {
			log.Fatal("Error loading env file:", err)
		}
	}

	cfg := &Config{
		AppEnv: getEnv("APP_ENV", "development"),
		Port: getEnv("PORT", "3000"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		JWTSecret: getEnv("JWT_SECRET", "changeme"),
	}

	return cfg
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	
	return defaultVal
}

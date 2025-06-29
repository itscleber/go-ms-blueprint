package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvAndLogger() string {
	SetupLogger()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatalf("SERVICE_NAME is not set")
	}

	return serviceName
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

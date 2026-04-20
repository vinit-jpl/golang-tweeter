package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	DBUrlMigration string
	SecretJwt      string

	DBHost     string
	DBUser     string
	DBName     string
	DBPassword string
	DBPort     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, fmt.Errorf("failed to load .env file")
	}

	log.Println("config loaded")

	return &Config{
		Port:           os.Getenv("PORT"),
		DBUrlMigration: os.Getenv("DATABASE_URL"),
		SecretJwt:      os.Getenv("SECRET_JWT"),
		DBHost:         os.Getenv("DB_HOST"),
		DBUser:         os.Getenv("DB_USER"),
		DBName:         os.Getenv("DB_NAME"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBPort:         os.Getenv("DB_PORT"),
	}, nil
}

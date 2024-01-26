package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Error("Error loading .env file")
	}
}

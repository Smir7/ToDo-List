package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"simple-service/internal/api"
	"simple-service/internal/config"
)

func main() {
	// Загрузка конфигурации
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	app := fiber.New()
	api.SetupRoutes(app)
	log.Printf("Server is running on port %s", config.AppConfig.Port)
	if err := app.Listen(":" + config.AppConfig.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

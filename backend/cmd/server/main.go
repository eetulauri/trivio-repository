package main

import (
	"log"
	"trivia-game/internal/config"
	"trivia-game/internal/handlers"
	"trivia-game/pkg/gemini"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load configuration
	cfg := config.New()

	// Initialize Gemini client
	geminiClient, err := gemini.New(cfg.GeminiAPIKey)
	if err != nil {
		log.Fatalf("Failed to initialize Gemini client: %v", err)
	}
	defer geminiClient.Close()

	// Create new Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Trivia Game API",
	})

	// Add middleware
	app.Use(cors.New())   // Enable CORS
	app.Use(logger.New()) // Add logging

	// Initialize handlers
	triviaHandler := handlers.NewTriviaHandler(geminiClient)

	// Setup routes
	api := app.Group("/api")
	{
		// Health check route
		api.Get("/health", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"status":  "ok",
				"message": "Server is running",
			})
		})

		// Trivia routes
		api.Get("/question", triviaHandler.GetQuestion)
		api.Post("/check-answer", triviaHandler.CheckAnswer)
	}

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

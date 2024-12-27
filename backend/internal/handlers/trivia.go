package handlers

import (
	"log"
	"trivia-game/pkg/gemini"

	"github.com/gofiber/fiber/v2"
)

// TriviaHandler handles trivia-related requests
type TriviaHandler struct {
	geminiClient *gemini.Client
}

// NewTriviaHandler creates a new TriviaHandler
func NewTriviaHandler(geminiClient *gemini.Client) *TriviaHandler {
	return &TriviaHandler{
		geminiClient: geminiClient,
	}
}

// GetQuestion generates and returns a new trivia question
func (h *TriviaHandler) GetQuestion(c *fiber.Ctx) error {
	if h.geminiClient == nil {
		log.Println("Error: Gemini client is nil")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gemini client not initialized properly",
		})
	}

	question, err := h.geminiClient.GenerateQuestion()
	if err != nil {
		log.Printf("Error generating question: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"question": question,
	})
}

// CheckAnswer checks if the provided answer is correct
func (h *TriviaHandler) CheckAnswer(c *fiber.Ctx) error {
	if h.geminiClient == nil {
		log.Println("Error: Gemini client is nil")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gemini client not initialized properly",
		})
	}

	var request struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	}

	if err := c.BodyParser(&request); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	response, err := h.geminiClient.CheckAnswer(request.Question, request.Answer)
	if err != nil {
		log.Printf("Error checking answer: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"correct":  response.Correct,
		"feedback": response.Feedback,
		"tidbit":   response.Tidbit,
	})
}

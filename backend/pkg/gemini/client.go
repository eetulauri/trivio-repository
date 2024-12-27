package gemini

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// Client represents a Gemini API client
type Client struct {
	apiKey string
	client *genai.Client
	model  *genai.GenerativeModel
}

// New creates a new Gemini API client
func New(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("Gemini API key not set")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("error creating Gemini client: %v", err)
	}

	model := client.GenerativeModel("gemini-pro")
	model.SetTemperature(0.7)     // Set creativity level
	model.SetTopK(40)             // Consider top 40 tokens
	model.SetTopP(0.95)           // Nucleus sampling
	model.SetMaxOutputTokens(100) // Limit response length

	return &Client{
		apiKey: apiKey,
		client: client,
		model:  model,
	}, nil
}

// Close closes the Gemini client
func (c *Client) Close() {
	if c.client != nil {
		c.client.Close()
	}
}

// GenerateQuestion generates a random trivia question
func (c *Client) GenerateQuestion() (string, error) {
	ctx := context.Background()

	prompt := "Generate a random trivia question. Format: Just the question text only, no answer."
	resp, err := c.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("error generating question: %v", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response from API")
	}

	question, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
	if !ok {
		return "", fmt.Errorf("unexpected response format")
	}

	return string(question), nil
}

// CheckAnswer checks if the provided answer is correct for the given question
func (c *Client) CheckAnswer(question, userAnswer string) (bool, error) {
	ctx := context.Background()

	prompt := fmt.Sprintf("Question: %s\nUser's Answer: %s\nIs this answer correct? Just respond with 'true' or 'false'.",
		question, userAnswer)

	resp, err := c.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return false, fmt.Errorf("error checking answer: %v", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return false, fmt.Errorf("no response from API")
	}

	response, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
	if !ok {
		return false, fmt.Errorf("unexpected response format")
	}

	return string(response) == "true", nil
}

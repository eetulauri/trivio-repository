package gemini

import (
	"context"
	"encoding/json"
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

// AnswerResponse represents the response for an answer check
type AnswerResponse struct {
	Correct  bool   `json:"correct"`
	Tidbit   string `json:"tidbit"`
	Feedback string `json:"feedback"`
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
	model.SetMaxOutputTokens(200) // Increased for tidbits

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

	prompt := `Generate a random trivia question. The question should be interesting and educational.
Format: Just the question text only, no answer or additional information.`

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

// CheckAnswer checks if the provided answer is correct and returns educational information
func (c *Client) CheckAnswer(question, userAnswer string) (*AnswerResponse, error) {
	ctx := context.Background()

	prompt := fmt.Sprintf(`Analyze this trivia question and answer:
Question: %s
User's Answer: %s

Provide feedback and educational information in this exact JSON format:
{
  "correct": true/false,
  "feedback": "Brief feedback about the answer (1 sentence)",
  "tidbit": "An interesting fact related to the topic (1-2 sentences)"
}`, question, userAnswer)

	resp, err := c.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("error checking answer: %v", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("no response from API")
	}

	response, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	// Parse the JSON response
	var result AnswerResponse
	if err := json.Unmarshal([]byte(response), &result); err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	return &result, nil
}

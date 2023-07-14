package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/NumexaHQ/captainCache/cache"
	"github.com/gofiber/fiber/v2"
	openai "github.com/sashabaranov/go-openai"
)

// HandleRequest handles the incoming request
func HandleRequest(c *fiber.Ctx) error {
	apiKey := c.Query("api_key")
	prompt := c.Query("prompt")

	if apiKey == "" || prompt == "" {
		return c.Status(http.StatusBadRequest).SendString("API key and prompt are required")
	}
	// Check if the prompt is cached in Redis
	cachedResponse, err := cache.GetFromCache(prompt)
	if err == nil {
		// Serve the cached response
		return c.SendString(fmt.Sprintf("Cached Response:\n%s", cachedResponse))
	}

	client := openai.NewClient(apiKey)

	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to call OpenAI API")
	}

	// Cache the response in Redis for future use
	err = cache.StoreInCache(prompt, response.Choices[0].Message.Content)
	if err != nil {
		log.Printf("Failed to cache response for prompt '%s': %v", prompt, err)
	}
	// Serve the API response
	return c.SendString(fmt.Sprintf("API Response:\n%s", response.Choices[0].Message.Content))
}

// app/cache/cache.go
package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func init() {
	// Connect to Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Enter your Redis password if applicable
		DB:       0,  // Use default Redis database
	})
}

// GetFromCache retrieves the cached response from Redis based on the prompt
func GetFromCache(prompt string) (string, error) {
	cachedResponse, err := redisClient.Get(context.Background(), prompt).Result()
	if err != nil {
		return "", err
	}
	return cachedResponse, nil
}

// StoreInCache stores the response in Redis based on the prompt
func StoreInCache(prompt string, response string) error {
	err := redisClient.Set(context.Background(), prompt, response, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

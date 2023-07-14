// main.go
package main

import (
	"log"

	"github.com/NumexaHQ/captainCache/routes"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var redisClient *redis.Client

func main() {
	// Connect to Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Enter your Redis password if applicable
		DB:       0,  // Use default Redis database
	})

	// Create a new Fiber app
	app := fiber.New()

	// Use logger middleware
	app.Use(logger.New())

	// Initialize routes
	routes.Setup(app)

	// Start the server
	log.Fatal(app.Listen(":8080"))
}

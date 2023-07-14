// routes/routes.go
package routes

import (
	"github.com/NumexaHQ/captainCache/handlers"
	"github.com/gofiber/fiber/v2"
)

// Setup initializes the routes
func Setup(app *fiber.App) {
	app.Get("/", handlers.HandleRequest)
}

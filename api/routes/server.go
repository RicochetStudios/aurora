package routes

import (
	"ricochet/aurora/api/services"

	"github.com/gofiber/fiber/v2"
)

// ServerRouter is the router for all server methods.
func ServerRouter(app fiber.Router) {
	// Get server details.
	app.Post("/server", services.GetServer())

	// Update server details.
	app.Put("/server", services.UpdateServer())

	// Remove server.
	app.Delete("/server", services.RemoveServer())
}

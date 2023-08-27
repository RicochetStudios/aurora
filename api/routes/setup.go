package routes

import (
	"github.com/RicochetStudios/aurora/api/services"

	"github.com/gofiber/fiber/v2"
)

// SetupRouter is the router for all setup methods.
func SetupRouter(app fiber.Router) {
	// Run setup.
	app.Post("/setup", services.Setup())
}

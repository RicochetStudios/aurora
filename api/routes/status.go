package routes

import (
	"github.com/RicochetStudios/aurora/api/services"

	"github.com/gofiber/fiber/v2"
)

// StatusRouter is the router for all setup methods.
func StatusRouter(app fiber.Router) {
	// Run hit test.
	app.Get("/", services.HitTest())
}

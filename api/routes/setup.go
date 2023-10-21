package routes

import (
	"github.com/RicochetStudios/aurora/api/middleware"
	"github.com/RicochetStudios/aurora/api/models"
	"github.com/RicochetStudios/aurora/api/services"

	"github.com/gofiber/fiber/v2"
)

// SetupRouter is the router for all setup methods.
func SetupRouter(app fiber.Router) {
	// Run setup.
	app.Post("/setup", func(c *fiber.Ctx) error {return middleware.ValidatePayload(models.GameServer{}, c)}, middleware.Auth, services.Setup())
}

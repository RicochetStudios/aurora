package routes

import (
	"github.com/RicochetStudios/aurora/api/middleware"
	"github.com/RicochetStudios/aurora/api/models"
	"github.com/RicochetStudios/aurora/api/services"

	"github.com/gofiber/fiber/v2"
)

// ServerRouter is the router for all server methods.
func ServerRouter(app fiber.Router) {
	// Get server details.
	app.Post("/server", middleware.Auth, services.GetServer())

	// Update server details.
	app.Put("/server", func(c *fiber.Ctx) error {return middleware.ValidatePayload(models.GameServer{}, c)}, middleware.Auth, services.UpdateServer())

	// Remove server.
	app.Delete("/server", middleware.Auth, services.RemoveServer())
}

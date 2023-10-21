package routes

import (
	"github.com/RicochetStudios/aurora/api/middleware"
	"github.com/RicochetStudios/aurora/api/services"

	"github.com/gofiber/fiber/v2"
)

// AuthRouter is the router for all auth methods.
func AuthRouter(app fiber.Router) {
	// set user auth.
	app.Post("/setAuthUser", middleware.Auth, services.SetAuthUser())

	// Verify user jwt token.
	app.Post("/verifyAuthUser", middleware.Auth, services.VerifyAuthUser())
}

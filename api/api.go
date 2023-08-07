package api

import (
	"log"
	"ricochet/aurora/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start() {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")

	// Hit test
	api.Get("/", routes.HitTest)

	// Get server details.
	api.Post("/server", routes.GetServer)

	// Update server details
	api.Put("/server", routes.UpdateServer)

	// Update server details
	api.Delete("/server", routes.RemoveServer)

	// update date into firebase (TESTING)
	api.Post("/server/firebase", routes.UpdateServerFromFirebase)

	// get data from firebase (TESTING)
	api.Get("/server/firebase", routes.GetServerFromFirebase)

	// Start the API.
	log.Fatal(app.Listen(":6969"))
}

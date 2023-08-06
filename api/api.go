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

	// update server details
	api.Put("/server", routes.UpdateServer)

	// Start the API.
	log.Fatal(app.Listen(":6969"))
}



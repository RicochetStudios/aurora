package api

import (
	"log"

	"github.com/RicochetStudios/aurora/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start() {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")

	// Run the status router.
	routes.StatusRouter(api)

	// Run the setup router.
	routes.SetupRouter(api)

	// Run the server router.
	routes.ServerRouter(api)

	// Start the API.
	log.Fatal(app.Listen(":6969"))
}

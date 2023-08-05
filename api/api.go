package api

import (
	"log"
	"ricochet/aurora/types"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// RunApi runs the server api.
func RunApi() {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")

	// Test handler
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("App running")
	})

	// Get server details.
	api.Get("/server", func(c *fiber.Ctx) error {
		server, _ := GetServer()
		return c.JSON(server)
	})

	log.Fatal(app.Listen(":5000"))
}

func GetServer() (types.Server, error) {
	// Temporarily create mock server response.
	mockServer := types.Server{
		ID:   "00000001",
		Size: "xs",
		Game: types.Game{
			Name:      "minecraft_java",
			Modloader: "vanilla",
		},
		Network: types.Network{
			Type: "private",
		},
	}

	return mockServer, nil
}

package api

import (
	"log"
	"ricochet/aurora/docker"
	"ricochet/aurora/types"

	"github.com/docker/go-connections/nat"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// RunApi runs the server api.
func RunApi() {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")

	// Test handler
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("App running")
	})

	// Get server details.
	api.Get("/server", func(ctx *fiber.Ctx) error {
		server, _ := GetServer()
		return ctx.JSON(server)
	})

	// Creates or updates a server.
	api.Put("/server", func(ctx *fiber.Ctx) error {
		server := new(types.Server)

		// Check for errors in body.
		if err := ctx.BodyParser(server); err != nil {
			return err
		}

		err := UpdateServer(ctx, server)
		if err != nil {
			return err
		}

		return ctx.JSON(server)
	})

	log.Fatal(app.Listen(":5000"))
}

// UpdateServer creates or updates a server.
func UpdateServer(ctx *fiber.Ctx, server *types.Server) error {
	// Create container environment variables.
	env, err := docker.NewContainerEnvVar("name", "value")
	if err != nil {
		log.Fatal(err)
	}
	// Create container environment ports.
	port, err := nat.NewPort("tcp", "8080")
	if err != nil {
		log.Fatal(err)
	}

	// Create container config.
	config, err := docker.NewContainerConfig(
		server.ID,
		"nginx",
		nat.PortSet{port: struct{}{}},
		[]string{"/data:/data"},
		[]string{env},
	)
	if err != nil {
		log.Fatal(err)
	}

	container, err := docker.RunServer(ctx, config)

	return err
}

// GetServer gets details about the currently configured game server instance.
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

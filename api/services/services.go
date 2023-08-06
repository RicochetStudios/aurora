package services

import (
	"log"
	"ricochet/aurora/types"

	"github.com/docker/go-connections/nat"
	"github.com/gofiber/fiber/v2"

	"ricochet/aurora/docker"
)

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

func NewServer() *types.Server {
	return new(types.Server)

}

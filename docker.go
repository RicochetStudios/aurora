package main

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// A volume for persistent data.
type ServerVolume struct {
	DataPath  string
	MountPath string
}

// Environment variable to expose to the application.
type ServerEnvVar struct {
	Name  string
	Value string
}

// Input for the server config.
type ServerConfig struct {
	Image   string
	Ports   []nat.Port
	Volumes []ServerVolume
	Envs    []ServerEnvVar
}

// Creates a server container and starts it. Similar to `docker run`.
func RunServer(ctx context.Context, cli *client.Client) (container.CreateResponse, error) {
	// Create the container.
	response, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        "nginx",
		ExposedPorts: nat.PortSet{"8080": struct{}{}},
	}, &container.HostConfig{
		PortBindings: map[nat.Port][]nat.PortBinding{nat.Port("8080"): {{HostIP: "127.0.0.1", HostPort: "8080"}}},
	}, nil, nil, "nginx-go-cli")
	if err != nil {
		return response, err
	}

	// Start the container.
	if err := cli.ContainerStart(ctx, response.ID, types.ContainerStartOptions{}); err != nil {
		return response, err
	}

	return response, nil
}

// Stops and removes a server container.
func RemoveServer(ctx context.Context, cli *client.Client, containerID string) error {
	// Stop and delete the container and volumes.
	if err := cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   false,
		Force:         true,
	}); err != nil {
		return err
	}

	// Remove unused data.
	_, err := cli.ContainersPrune(ctx, filters.Args{})
	if err != nil {
		return err
	}

	return nil
}

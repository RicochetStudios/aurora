package docker

import (
	"context"
	"errors"
	"io"
	"os"
	"regexp"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// NewContainerEnvVar creates a new instance of ContainerEnvVar given a name and value.
func NewContainerEnvVar(name, value string) (string, error) {
	match, err := regexp.MatchString(`^[a-zA-Z_][a-zA-Z0-9_]*$`, name)
	if err != nil {
		return "", err
	} else if !match {
		// If the name provided is not a valid environment variable name, error.
		return "", errors.New("environment variable name '" + name + "' is not valid")
	} else {
		return (name + "=" + value), nil
	}
}

// ContainerConfig is a set of configurations to pass to the docker engine to create the server container.
type ContainerConfig struct {
	Name         string
	Image        string
	ExposedPorts nat.PortSet
	Binds        []string
	Env          []string
}

// NewContainerConfig creates a new ContainerConfig.
func NewContainerConfig(name, image string, ports nat.PortSet, binds []string, env []string) (ContainerConfig, error) {
	containerConfig := ContainerConfig{
		Name:         name,
		Image:        image,
		ExposedPorts: ports,
		Binds:        binds,
		Env:          env,
	}

	return containerConfig, nil
}

// Creates a server container and starts it. Similar to `docker run`.
func RunServer(ctx context.Context, cli *client.Client, config ContainerConfig) (container.CreateResponse, error) {
	// Pull the image.
	out, err := cli.ImagePull(ctx, config.Image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	// Output the download status.
	io.Copy(os.Stdout, out)

	// Create the container.
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        config.Image,
		ExposedPorts: config.ExposedPorts,
	}, &container.HostConfig{
		// Binds work the way that mounts would normally.
		Binds: config.Binds,
		// Not sure if we need host bindings yet.
		// PortBindings: map[nat.Port][]nat.PortBinding{nat.Port("8080"): {{HostIP: "127.0.0.1", HostPort: "8080"}}},
	}, nil, nil, config.Name)
	if err != nil {
		return resp, err
	}

	// Start the container.
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return resp, err
	}

	return resp, nil
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

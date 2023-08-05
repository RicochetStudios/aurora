package docker

import (
	"context"
	"errors"
	"io"
	"os"
	"regexp"

	dockerTypes "github.com/docker/docker/api/types"
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

// RunServer creates a server container and starts it. Similar to `docker run`.
func RunServer(ctx context.Context, config ContainerConfig) (container.CreateResponse, error) {
	// Constructs the client object.
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return container.CreateResponse{}, err
	}
	defer cli.Close()

	// Pull the image.
	out, err := cli.ImagePull(ctx, config.Image, dockerTypes.ImagePullOptions{})
	if err != nil {
		return container.CreateResponse{}, err
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
	if err := cli.ContainerStart(ctx, resp.ID, dockerTypes.ContainerStartOptions{}); err != nil {
		return resp, err
	}

	return resp, nil
}

// RemoveServer stops and removes a server container.
func RemoveServer(ctx context.Context, containerID string) error {
	// Constructs the client object.
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	// Stop and delete the container and volumes.
	if err := cli.ContainerRemove(ctx, containerID, dockerTypes.ContainerRemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   false,
		Force:         true,
	}); err != nil {
		return err
	}

	// Remove unused data.
	if _, err := cli.ContainersPrune(ctx, filters.Args{}); err != nil {
		return err
	}

	return nil
}

// // GetServer gets details about the currently configured game server instance.
// func GetServer(ctx context.Context, cli *client.Client, containerID string) (types.Server, error) {
// 	// containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
// 	// 	Filters: filters.NewArgs(filters.Arg("id", containerID)),
// 	// })
// 	// if err != nil {
// 	// 	return resp, err
// 	// }

// 	// for _, container := range containers {
// 	// 	fmt.Printf("%s %s\n", container.ID[:10], container.Image)
// 	// }
// }

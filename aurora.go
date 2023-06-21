package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {
	// Constructs the client object.
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	// Run a mongo container.
	RunServer(ctx, cli)

	// Gets containers that are actively running.
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	// Outputs running containers.
	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}

func RunServer(ctx context.Context, cli *client.Client) {
	// Create the container.
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        "nginx",
		ExposedPorts: nat.PortSet{"8080": struct{}{}},
	}, &container.HostConfig{
		PortBindings: map[nat.Port][]nat.PortBinding{nat.Port("8080"): {{HostIP: "127.0.0.1", HostPort: "8080"}}},
	}, nil, nil, "nginx-go-cli")
	if err != nil {
		panic(err)
	}

	// Start the container.
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func main() {
	// Constructs the client object.
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Create container config.
	config, err := NewContainerConfig(
		"my-unique-id",
		"nginx",
		[]ServerPort{NewServerPort("tcp", "8080")},
		[]ServerMount{},
		[]ServerEnvVar{},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Run an nginx container.
	server, err := RunServer(ctx, cli, config)
	if err != nil {
		log.Fatal(err)
	}

	// Gets containers that are actively running.
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		Filters: filters.NewArgs(filters.Arg("id", server.ID)),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Outputs running containers.
	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}

	// Removes the container.
	err = RemoveServer(ctx, cli, server.ID)
	if err != nil {
		log.Fatal(err)
	}
}

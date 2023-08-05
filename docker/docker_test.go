package docker

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/google/go-cmp/cmp"
)

// // TestNewServerPort calls NewServerPort with a protocol and port,
// // checking for a valid ServerPort type to be returned.
// func TestNewServerPort(t *testing.T) {
// 	var want ServerPort = ServerPort{Protocol: "tcp", Port: "8080"}
// 	got := NewServerPort("tcp", "8080")
// 	if want != got {
// 		t.Fatalf(`NewServerPort("tcp", "8080") = %q, want match for %#q`, got, want)
// 	}
// }

// TestNewContainerEnvVar calls NewContainerEnvVar with a name and value,
// checking for a valid string in return.
func TestNewContainerEnvVar(t *testing.T) {
	var want string = "EULA=TRUE"
	got, err := NewContainerEnvVar("EULA", "TRUE")
	if want != got || err != nil {
		t.Fatalf(`NewServerVolume("EULA", "TRUE") = %q, %v, want match for %#q, nil`, got, err, want)
	}
}

// TestNewContainerEnvVarInvalidName calls NewContainerEnvVar with an invalid name,
// checking for an error in return.
func TestNewContainerEnvVarInvalidName(t *testing.T) {
	// Add invalid characters to the env var name.
	_, err := NewContainerEnvVar("$EULA!", "TRUE")
	if err == nil {
		t.Fatalf(`NewContainerEnvVar("$EULA!", "TRUE") expected a name invalid error while testing, got %T`, err)
	}
}

// TestNewContainerConfig calls NewContainerConfig with correct inputs,
// checking for a valid ContainerConfig in return.
func TestNewContainerConfig(t *testing.T) {
	var want ContainerConfig = ContainerConfig{
		"my-unique-id",
		"nginx",
		nat.PortSet{"8080/tcp": struct{}{}},
		[]string{"/data:/data"},
		[]string{"name=value"},
	}

	got, err := NewContainerConfig(
		"my-unique-id",
		"nginx",
		nat.PortSet{"8080/tcp": struct{}{}},
		[]string{"/data:/data"},
		[]string{"name=value"},
	)

	if err != nil {
		t.Fatalf("NewContainerConfig() returned an error: \n%v", err)
	}
	// Use cmp for more complex types.
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("NewContainerConfig() mismatch (-want +got):\n%s", diff)
	}
}

// cleaupAllContainers removes all running containers, volumes and data.
func cleaupAllContainers(ctx context.Context, cli *client.Client) error {
	// Get all containers.
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}

	// Stop and remove containers and volumes.
	for _, instance := range containers {
		fmt.Print("Stopping container ", instance.ID[:10], "... ")
		if err := cli.ContainerRemove(ctx, instance.ID, types.ContainerRemoveOptions{
			RemoveVolumes: true,
			RemoveLinks:   false,
			Force:         true,
		}); err != nil {
			return err
		}
		fmt.Println("Success")
	}

	// Remove unused data.
	if _, err := cli.ContainersPrune(ctx, filters.Args{}); err != nil {
		return err
	}

	return nil
}

// TestRunServer calls RunServer with ContainerConfig,
// checking for a valid CreateResponse in return.
func TestRunServer(t *testing.T) {
	ctx := context.Background()
	// Constructs the client object.
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()

	// Run a test container.
	got, err := RunServer(ctx, cli, ContainerConfig{
		"my-unique-id",
		"nginx",
		nat.PortSet{"8080/tcp": struct{}{}},
		[]string{"/data:/data"},
		[]string{"name=value"},
	})

	if err != nil {
		t.Fatalf("RunServer() returned an error: \n%v", err)
	}
	// Use cmp for more complex types.
	if cmp.Equal(got, container.CreateResponse{}) {
		t.Fatalf(`RunServer() returned an empty response.`)
	}
	if len(got.Warnings) > 0 {
		t.Fatalf(`RunServer() returned warnings in the response.`)
	}

	// Cleanup any remaining containers, volumes and data.
	t.Cleanup(func() {
		if cleanupErr := cleaupAllContainers(ctx, cli); cleanupErr != nil {
			t.Fatalf("TestRunServer() error cleaning up:\n%v", cleanupErr)
		}
	})
}

// TestRemoveServer calls RemoveServer,
// checking for no errors in return.
func TestRemoveServer(t *testing.T) {
	ctx := context.Background()
	// Constructs the client object.
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()

	// Run a test container.
	server, err := RunServer(ctx, cli, ContainerConfig{
		"my-unique-id",
		"nginx",
		nat.PortSet{"8080/tcp": struct{}{}},
		[]string{"/data:/data"},
		[]string{"name=value"},
	})
	if err != nil {
		t.Fatalf("RunServer() returned an error: \n%v", err)
	}

	// Stop the container.
	if err := RemoveServer(ctx, cli, server.ID); err != nil {
		t.Fatalf("RemoveServer() returned an error: \n%v", err)
	}

	// Cleanup any remaining containers, volumes and data.
	t.Cleanup(func() {
		if cleanupErr := cleaupAllContainers(ctx, cli); cleanupErr != nil {
			t.Fatalf("TestRunServer() error cleaning up:\n%v", cleanupErr)
		}
	})
}

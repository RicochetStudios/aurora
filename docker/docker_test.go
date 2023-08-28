package docker

import (
	"context"
	"fmt"
	"testing"

	"github.com/RicochetStudios/aurora/schema"
	"github.com/RicochetStudios/aurora/types"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/google/go-cmp/cmp"
)

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

// cleaupAllContainers removes all running containers, volumes and data.
func cleaupAllContainers(ctx context.Context, cli *client.Client) error {
	// Get all containers.
	containers, err := cli.ContainerList(ctx, dockerTypes.ContainerListOptions{})
	if err != nil {
		return fmt.Errorf("cleaupAllContainers() error getting list of containers: %v", err)
	}

	// Stop and remove containers and volumes.
	for _, instance := range containers {
		fmt.Print("Stopping container ", instance.ID[:10], "... ")
		if err := cli.ContainerRemove(ctx, instance.ID, dockerTypes.ContainerRemoveOptions{
			RemoveVolumes: true,
			RemoveLinks:   false,
			Force:         true,
		}); err != nil {
			return fmt.Errorf("cleaupAllContainers() error removing container with id: %v:\n%v", instance.ID, err)
		}
		fmt.Println("Success")
	}

	// Remove unused data.
	if _, err := cli.ContainersPrune(ctx, filters.Args{}); err != nil {
		return fmt.Errorf("cleaupAllContainers() error pruning containers: %v", err)
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

	// Cleanup any remaining containers, volumes and data at the end of the test.
	t.Cleanup(func() {
		if cleanupErr := cleaupAllContainers(ctx, cli); cleanupErr != nil {
			t.Fatalf("TestRunServer() error cleaning up:\n%v", cleanupErr)
		}
	})

	// Run a test container.
	got, err := RunServer(ctx, ContainerConfig{
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

	// Cleanup any remaining containers, volumes and data at the end of the test.
	t.Cleanup(func() {
		if cleanupErr := cleaupAllContainers(ctx, cli); cleanupErr != nil {
			t.Fatalf("TestRunServer() error cleaning up:\n%v", cleanupErr)
		}
	})

	// Run a test container.
	if _, err := RunServer(ctx, ContainerConfig{
		"my-unique-id",
		"nginx",
		nat.PortSet{"8080/tcp": struct{}{}},
		[]string{"/data:/data"},
		[]string{"name=value"},
	}); err != nil {
		t.Fatalf("RunServer() returned an error: \n%v", err)
	}

	// Stop the container.
	if err := RemoveServer(ctx); err != nil {
		t.Fatalf("RemoveServer() returned an error: \n%v", err)
	}
}

// TestNewContainerConfigFromSchema calls NewContainerConfigFromSchema with correct inputs,
// checking for a valid ContainerConfig in return.
func TestNewContainerConfigFromSchema(t *testing.T) {
	// Create the schema.
	var schema schema.Schema = schema.Schema{
		Name:  "minecraft_java",
		Image: "itzg/minecraft-server:latest",
		URL:   "https://github.com/itzg/docker-minecraft-server",
		Ratio: "1-2",
		Sizes: map[string]schema.Size{
			"xs": {
				Resources: schema.Resources{CPU: "1000m", Memory: "2000Mi"},
				Players:   8,
			},
			"s": {
				Resources: schema.Resources{CPU: "1500m", Memory: "4000Mi"},
				Players:   16,
			},
			"m": {
				Resources: schema.Resources{CPU: "2000m", Memory: "8000Mi"},
				Players:   32,
			},
			"l": {
				Resources: schema.Resources{CPU: "3000m", Memory: "16000Mi"},
				Players:   64,
			},
			"xl": {
				Resources: schema.Resources{CPU: "4000m", Memory: "32000Mi"},
				Players:   128,
			},
		},
		Network: []schema.Network{
			{Name: "game", Port: 25565, Protocol: "tcp"},
		},
		Settings: []schema.Setting{
			{Name: "EULA", Value: "TRUE"},
			{Name: "TYPE", Value: "{{ .modloader }}"},
			{Name: "MAX_PLAYERS", Value: "{{ .players }}"},
			{Name: "MOTD", Value: "{{ .name }}"},
		},
		Volumes: []schema.Volume{
			{
				Name:  "data",
				Path:  "/data",
				Class: "classic",
				Size:  "10Gi",
			},
		},
		Probes: schema.Probes{
			Command: []string{"mc-health"},
			StartupProbe: schema.Probe{
				FailureThreshold: 30,
				PeriodSeconds:    10,
			},
			ReadynessProbe: schema.Probe{
				InitialDelaySeconds: 30,
				PeriodSeconds:       5,
				FailureThreshold:    20,
				SuccessThreshold:    3,
				TimeoutSeconds:      1,
			},
			LivenessProbe: schema.Probe{
				InitialDelaySeconds: 30,
				PeriodSeconds:       5,
				FailureThreshold:    20,
				SuccessThreshold:    3,
				TimeoutSeconds:      1,
			},
		},
	}

	var server types.Server = types.Server{
		Name: "mytest",
		Size: "xs",
		Game: types.Game{
			Name:      "minecraft_java",
			Modloader: "vanilla",
		},
		Network: types.Network{
			Type: "private",
		},
	}

	var want ContainerConfig = ContainerConfig{
		"my-unique-id",
		"itzg/minecraft-server:latest",
		nat.PortSet{"25565/tcp": struct{}{}},
		[]string{"/data:/data"},
		[]string{
			"EULA=TRUE",
			"TYPE=vanilla",
			"MAX_PLAYERS=8",
			"MOTD=mytest",
		},
	}

	got, err := NewContainerConfig("my-unique-id", schema, server)

	if err != nil {
		t.Fatalf("NewContainerConfigFromSchema() returned an error: \n%v", err)
	}
	// Use cmp for more complex types.
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("NewContainerConfigFromSchema() mismatch (-want +got):\n%s", diff)
	}
}

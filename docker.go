package main

import (
	"context"
	"errors"
	"regexp"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// ServerPort is a set of port number and protocol to expose from the server.
type ServerPort struct {
	Port     string
	Protocol string
}

// NewServerPort creates a new instance of ServerPort given a protocol and port number.
func NewServerPort(proto, port string) ServerPort {
	return ServerPort{
		Port:     port,
		Protocol: proto,
	}
}

// A volume for persistent data.
type ServerMount string

func NewServerMount(target string) (ServerMount, error) {
	match, err := regexp.MatchString(`^/.*$`, target)
	if err != nil {
		return ServerMount(""), err
	} else if !match {
		// If the target provided is not a valid unix path, error.
		return ServerMount(""), errors.New("mount target '" + target + "' is not a valid path")
	} else {
		return ServerMount(target), nil
	}
}

// ServerEnvVar is a key value pair of an environment variable to provide to the server.
type ServerEnvVar struct {
	Name  string
	Value string
}

// NewServerEnvVar creates a new instance of ServerEnvVar given a name and value.
func NewServerEnvVar(name, value string) (ServerEnvVar, error) {
	match, err := regexp.MatchString(`^[a-zA-Z_][a-zA-Z0-9_]*$`, name)
	if err != nil {
		return ServerEnvVar{}, err
	} else if !match {
		// If the name provided is not a valid environment variable name, error.
		return ServerEnvVar{}, errors.New("environment variable name '" + name + "' is not valid")
	} else {
		return ServerEnvVar{
			Name:  name,
			Value: value,
		}, nil
	}
}

// ContainerConfig is a set of configurations to pass to the docker engine to create the server container.
type ContainerConfig struct {
	Name         string
	Image        string
	ExposedPorts nat.PortSet
	Mounts       []mount.Mount
	Env          []string
}

// NewContainerConfig creates a new ContainerConfig.
func NewContainerConfig(name, image string, ports []ServerPort, mounts []ServerMount, env []ServerEnvVar) (ContainerConfig, error) {
	containerPorts, err := NewContainerPorts(ports)
	if err != nil {
		return ContainerConfig{}, err
	}

	containerMounts := NewContainerMounts(mounts)

	containerEnv := NewContainerEnv(env)

	containerConfig := ContainerConfig{
		Name:         name,
		Image:        image,
		ExposedPorts: containerPorts,
		Mounts:       containerMounts,
		Env:          containerEnv,
	}

	return containerConfig, nil
}

// NewContainerPortSet creates a nat.PortSet from a list of ServerPorts.
func NewContainerPorts(serverPorts []ServerPort) (nat.PortSet, error) {
	var containerPorts nat.PortSet = nat.PortSet{}

	for _, portSet := range serverPorts {
		natPort, err := nat.NewPort(portSet.Protocol, portSet.Port)
		if err != nil {
			return nat.PortSet{}, err
		}

		containerPorts[natPort] = struct{}{}
	}

	return containerPorts, nil
}

// NewContainerMounts creates a list of container mounts from a list of server mounts.
func NewContainerMounts(serverMounts []ServerMount) []mount.Mount {
	var containerMounts []mount.Mount

	for _, mountTarget := range serverMounts {
		containerMount := mount.Mount{
			Source: string(mountTarget),
			Target: string(mountTarget),
		}
		containerMounts = append(containerMounts, containerMount)
	}

	return containerMounts
}

// NewContainerEnv creates a slice of container environment variables from a list of ServerEnvVars.
func NewContainerEnv(serverEnvVars []ServerEnvVar) []string {
	var containerEnv []string

	for _, envVar := range serverEnvVars {
		var containerVar string = envVar.Name + "=" + envVar.Value
		containerEnv = append(containerEnv, containerVar)
	}

	return containerEnv
}

// Creates a server container and starts it. Similar to `docker run`.
func RunServer(ctx context.Context, cli *client.Client, config ContainerConfig) (container.CreateResponse, error) {
	// Create the container.
	response, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        config.Image,
		ExposedPorts: config.ExposedPorts,
	}, &container.HostConfig{
		Mounts: config.Mounts,
		// Not sure if we need host bindings yet.
		// PortBindings: map[nat.Port][]nat.PortBinding{nat.Port("8080"): {{HostIP: "127.0.0.1", HostPort: "8080"}}},
	}, nil, nil, config.Name)
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

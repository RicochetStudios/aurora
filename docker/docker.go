package docker

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/RicochetStudios/aurora/schema"
	"github.com/RicochetStudios/aurora/types"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// templateRegex is a regular expression to validate templates.
const templateRegex string = `^{{ (?P<tpl>(\.\w+)*) }}$`

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

// templateValue takes a value and resolves its template if it is a template.
func templateValue(v string, g schema.Schema, s types.Server) string {
	// Template the env var if needed.
	re := regexp.MustCompile(templateRegex)
	if re.MatchString(v) {
		// Get the template to target.
		matches := re.FindStringSubmatch(v)
		tplIndex := re.SubexpIndex("tpl")
		tpl := matches[tplIndex]

		// Resolve the templates.
		switch tpl {
		case ".name":
			return s.Name
		case ".modloader":
			return s.Game.Modloader
		case ".players":
			return fmt.Sprint(g.Sizes[s.Size].Players)
		}
	}
	// If it is not a template, return the original value.
	return v
}

// NewContainerConfig creates a new ContainerConfig from a name, game schema and a server.
func NewContainerConfig(name string, gameSchema schema.Schema, server types.Server) (ContainerConfig, error) {
	// Create container environment ports.
	var portSet nat.PortSet = nat.PortSet{}
	for _, network := range gameSchema.Network {
		port, err := nat.NewPort(
			network.Protocol,
			fmt.Sprint(network.Port),
		)
		if err != nil {
			return ContainerConfig{}, err
		}
		portSet[port] = struct{}{}
	}

	// Create container bindings.
	var bindList []string = []string{}
	for _, volume := range gameSchema.Volumes {
		bindList = append(bindList, (volume.Path + ":" + volume.Path))
	}

	// Create container environment variables.
	var envList []string = []string{}
	for _, setting := range gameSchema.Settings {
		// Template environment variables if required.
		var sList [2]string = [2]string{setting.Name, setting.Value}
		for i, item := range sList {
			sList[i] = templateValue(item, gameSchema, server)
		}

		// Construct env vars.
		env, err := NewContainerEnvVar(
			sList[0],
			sList[1],
		)
		if err != nil {
			return ContainerConfig{}, err
		}
		envList = append(envList, env)
	}

	// Create container config.
	return ContainerConfig{
		Name:         name,
		Image:        gameSchema.Image,
		ExposedPorts: portSet,
		Binds:        bindList,
		Env:          envList,
	}, nil
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
func RemoveServer(ctx context.Context) error {
	// Constructs the client object.
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	// Get the latest container (the only one).
	containers, err := cli.ContainerList(ctx, dockerTypes.ContainerListOptions{
		Latest: true,
		All:    true,
	})
	if err != nil {
		return err
	}

	// Stop and delete the container and volumes.
	for _, cont := range containers {
		if err := cli.ContainerRemove(ctx, cont.ID, dockerTypes.ContainerRemoveOptions{
			RemoveVolumes: true,
			RemoveLinks:   false,
			Force:         true,
		}); err != nil {
			return err
		}
	}

	// Remove unused data.
	if _, err := cli.ContainersPrune(ctx, filters.Args{}); err != nil {
		return err
	}

	return nil
}

// // GetServer gets details about the currently configured game server instance.
// func GetServer(ctx context.Context, containerID string) (ContainerConfig, error) {
// 	// Constructs the client object.
// 	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
// 	if err != nil {
// 		return ContainerConfig{}, err
// 	}
// 	defer cli.Close()

// 	// Get the latest container (the only one).
// 	containers, err := cli.ContainerList(ctx, dockerTypes.ContainerListOptions{
// 		Latest: true,
// 		All:    true,
// 	})
// 	if err != nil {
// 		return ContainerConfig{}, err
// 	}

// 	var c dockerTypes.Container = containers[0]

// 	return ContainerConfig{
// 		Name:  c.Names[0],
// 		Image: c.Image,

// 		// Types are different.
// 		// ExposedPorts: c.Ports,
// 		// Binds:        c.Mounts,

// 		// Environment vars can't be returned.
// 		// Env:          envList,
// 	}, err
// }

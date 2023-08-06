package main

import (
	"ricochet/aurora/api"
)

func main() {
	// ctx := context.Background()

	// // Constructs the client object.
	// cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cli.Close()

	// Start the API.
	api.Start()

	// // // connect to firebase
	// app, err := db.FireBase()	

	// // connect to firebase realtime database
	// client, err := db.RealtimeDatabase()

	// // connect to firebase firestore
	// firestore, err := db.Firestore()

	// // Create container environment variables.
	// env, err := docker.NewContainerEnvVar("name", "value")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // Create container environment ports.
	// port, err := nat.NewPort("tcp", "8080")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Create container config.
	// config, err := docker.NewContainerConfig(
	// 	"my-unique-id",
	// 	"nginx",
	// 	nat.PortSet{port: struct{}{}},
	// 	[]string{"/data:/data"},
	// 	[]string{env},
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Run an nginx container.
	// server, err := docker.RunServer(ctx, config)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // // Gets containers that are actively running.
	// // containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
	// // 	Filters: filters.NewArgs(filters.Arg("id", server.ID)),
	// // })
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }

	// // // Outputs top 10 running containers.
	// // for _, container := range containers {
	// // 	fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	// // }

	// // Removes the container.
	// err = docker.RemoveServer(ctx, server.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

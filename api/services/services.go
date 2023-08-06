package services

import (
	"context"
	"fmt"
	"log"
	"ricochet/aurora/db"
	"ricochet/aurora/types"

	"github.com/docker/go-connections/nat"
	"github.com/gofiber/fiber/v2"

	"ricochet/aurora/docker"
)

// GetServer gets details about the currently configured game server instance.
func GetServer() (types.Server, error) {
	// Temporarily create mock server response.
	mockServer := types.Server{
		ID:   "00000001",
		Size: "xs",
		Game: types.Game{
			Name:      "minecraft_java",
			Modloader: "vanilla",
		},
		Network: types.Network{
			Type: "private",
		},
	}

	return mockServer, nil
}

// UpdateServer creates or updates a server.
func UpdateServer(ctx *fiber.Ctx, server *types.Server) error {
	// Create container environment variables.
	env, err := docker.NewContainerEnvVar("name", "value")
	if err != nil {
		log.Fatal(err)
	}

	// Create container environment ports.
	port, err := nat.NewPort("tcp", "8080")
	if err != nil {
		log.Fatal(err)
	}

	// Create container config.
	config, err := docker.NewContainerConfig(
		server.ID,
		"nginx",
		nat.PortSet{port: struct{}{}},
		[]string{"/data:/data"},
		[]string{env},
	)
	if err != nil {
		log.Fatal(err)
	}

	container, err := docker.RunServer(ctx, config)
	

	return err
}

func NewServer() *types.Server {
	return new(types.Server)
}

// test function to update data into firebase
func UpdateServerFromFirebase() {

	ctx := context.Background()
	
	firestore, err := db.Firestore()

	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	data := map[string]interface{}{
		"name": "Sam Woods",
	}

	// Set the value of 'NYC'.
	x, err := firestore.Collection("development").Doc("test").Set(ctx, data)

	if err != nil {
		log.Fatalln("Failed to update data: ", err)
	}
	

	fmt.Println("Updated data: ", x)
}

// test function to get data from firebase
func GetServerFromFirebase() interface{} {

	ctx := context.Background()
	
	firestore, err := db.Firestore()

	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	// Get the value from document 'test' in collection 'development'.
	dsnap, err := firestore.Collection("development").Doc("test").Get(ctx)

	if err != nil {
		log.Fatalln("Failed to get data: ", err)
	}

	data := dsnap.Data()
	
	return data
}
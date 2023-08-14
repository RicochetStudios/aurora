package services

import (
	"fmt"
	"log"
	"ricochet/aurora/config"
	"ricochet/aurora/db"
	"ricochet/aurora/schema"
	"ricochet/aurora/types"

	"github.com/gofiber/fiber/v2"

	"ricochet/aurora/docker"
)

// GetServer gets details about the currently configured game server instance.
func GetServer() types.Server {
	// Read the config file.
	config, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: \n%v", err)
	}
	var server types.Server = config.Server

	return server
}

// UpdateServer creates or updates a server.
func UpdateServer(ctx *fiber.Ctx, server types.Server) types.Server {
	// Get the game schema.
	schema, err := schema.GetSchema("minecraft_java")
	if err != nil {
		log.Fatalf("Error reading schema: \n%v", err)
	}

	// Create a container config.
	containerConfig, err := docker.NewContainerConfigFromSchema("my-unique-id", schema)
	if err != nil {
		log.Fatalf("Error creating container config: \n%v", err)
	}

	// Deploy and start the container.
	if _, err := docker.RunServer(ctx.Context(), containerConfig); err != nil {
		log.Fatalf("Error deploying container: \n%v", err)
	}

	// Update the local config.
	config, err := config.Update(config.Config{Server: server})
	if err != nil {
		log.Fatalf("Error updating local config: \n%v", err)
	}
	server = config.Server

	return server
}

// RemoveServer stops and deletes a server.
func RemoveServer(ctx *fiber.Ctx) {
	// Remove the container.
	if err := docker.RemoveServer(ctx.Context()); err != nil {
		log.Fatalf("Error removing container: \n%v", err)
	}

	// Remove the server details from the local config.
	_, err := config.Update(config.Config{Server: types.Server{
		ID:   "",
		Size: "",
		Game: types.Game{
			Name:      "",
			Modloader: "",
		},
		Network: types.Network{
			Type: "",
		},
	}})
	if err != nil {
		log.Fatalf("Error updating local config: \n%v", err)
	}
}

// test function to update data into firebase
func UpdateServerFromFirebase(ctx *fiber.Ctx) {
	firestore, err := db.Firestore(ctx.Context())

	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	data := map[string]interface{}{
		"name": "Sam Woods",
	}

	// Set the value of 'NYC'.
	x, err := firestore.Collection("development").Doc("test").Set(ctx.Context(), data)

	if err != nil {
		log.Fatalln("Failed to update data: ", err)
	}

	fmt.Println("Updated data: ", x)
}

// test function to get data from firebase
func GetServerFromFirebase(ctx *fiber.Ctx) interface{} {
	firestore, err := db.Firestore(ctx.Context())

	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	// Get the value from document 'test' in collection 'development'.
	dsnap, err := firestore.Collection("development").Doc("test").Get(ctx.Context())

	if err != nil {
		log.Fatalln("Failed to get data: ", err)
	}

	data := dsnap.Data()

	return data
}

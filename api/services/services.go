package services

import (
	"fmt"
	"log"
	"ricochet/aurora/config"
	"ricochet/aurora/db"
	"ricochet/aurora/docker"
	"ricochet/aurora/schema"
	"ricochet/aurora/types"

	"github.com/gofiber/fiber/v2"
)

// Setup performs initialisation steps to prepare the app to take following instructions.
func Setup(newConfig config.Config) config.Config {
	// Add or update the instance ID and cluster ID in the config.
	currentConfig, err := config.Update(newConfig)
	if err != nil {
		log.Fatalf("Error updating local config: \n%v", err)
	}

	return currentConfig
}

// GetServer gets details about the currently configured game server instance.
func GetServer(ctx *fiber.Ctx) types.Server {
	// Read the config file.
	config, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: \n%v", err)
	}

	// Read the current server configuration.
	server, err := db.GetServer(ctx.Context(), config.ID)
	if err != nil {
		log.Fatalf("Error reading server details from the database: \n%v", err)
	}

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

	// Read the config file.
	id, err := config.GetId()
	if err != nil {
		log.Fatalf("Error getting config id: \n%v", err)
	}

	// Add the status.
	server.Status = "running"

	// Create or update the current server configuration.
	server, err = db.SetServer(ctx.Context(), id, server)
	if err != nil {
		log.Fatalf("Error updating server details in the database: \n%v", err)
	}

	return server
}

// RemoveServer stops and deletes a server.
func RemoveServer(ctx *fiber.Ctx) {
	// Remove the container.
	if err := docker.RemoveServer(ctx.Context()); err != nil {
		log.Fatalf("Error removing container: \n%v", err)
	}

	// Read the config file.
	id, err := config.GetId()
	if err != nil {
		log.Fatalf("Error getting config id: \n%v", err)
	}

	// Delete the current server configuration.
	_, err = db.SetServer(ctx.Context(), id, types.Server{Status: "deallocated"})
	if err != nil {
		log.Fatalf("Error updating server details in the database: \n%v", err)
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

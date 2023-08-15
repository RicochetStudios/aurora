package services

import (
	"fmt"
	"log"
	"ricochet/aurora/config"
	"ricochet/aurora/db"

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

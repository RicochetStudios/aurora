package main

import (
	"fmt"
	"log"
	"os"
	"ricochet/aurora/api"

	"github.com/docker/docker/client"
)

func main() {
	// ctx := context.Background()

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(path)

	// Constructs the client object.
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// Start the API.
	api.Start()

	// connect to firebase
	// app, err := db.Firebase()

	// connect to firebase realtime database
	// client, err := db.RealtimeDatabase()

	// connect to firebase firestore
	// firestore, err := db.Firestore()
}

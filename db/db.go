package db

import (
	"context"
	"fmt"
	"log"
	"ricochet/aurora/types"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

// Call this function to get/initialise the firebase app
func Firebase() (*firebase.App, error) {
	ctx := context.Background()

	// configure database URL
	conf := &firebase.Config{
		DatabaseURL: "https://game-server-e2c56-default-rtdb.europe-west1.firebasedatabase.app",
	}

	// fetch service account key
	opt := option.WithCredentialsFile("config/firebase-config.json")

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	return app, err
}

// Call this function to get a client for the realtime database
func RealtimeDatabase() (*db.Client, error) {
	ctx := context.Background()

	app, err := Firebase()

	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("error in creating firebase DB client: ", err)
	}

	return client, err
}

// Call this function to get a client for the firestore database
func Firestore() (*firestore.Client, error) {
	ctx := context.Background()

	app, err := Firebase()

	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln("error in creating firebase DB client: ", err)
	}

	return client, err
}

// GetServer reads and returns a server document, given a server ID.
func GetServer(ctx context.Context, id string) (types.Server, error) {
	// Create the firestore client.
	client, err := Firestore()
	if err != nil {
		return types.Server{}, fmt.Errorf("error creating Firestore client:\n%v", err)
	}
	defer client.Close()

	// Read the full document from the database.
	// Temporarily hardcoding the collection, this needs to be changed to reflect the cluster it belongs to later.
	document, err := client.Collection("prod").Doc(id).Get(ctx)
	if err != nil {
		return types.Server{}, fmt.Errorf("error writing to Firestore database:\n%v", err)
	}

	// Convert the document into the server struct.
	var server types.Server
	if err := document.DataTo(&server); err != nil {
		return types.Server{}, fmt.Errorf("error converting Firestore document to types.Server struct:\n%v", err)
	}

	return server, nil
}

// SetServer creates and overwrites fields in the server document, given a Server.
func SetServer(ctx context.Context, server types.Server) (types.Server, error) {
	// Create the firestore client.
	client, err := Firestore()
	if err != nil {
		return types.Server{}, fmt.Errorf("error creating Firestore client:\n%v", err)
	}
	defer client.Close()

	// Write to the database, overwriting existing fields and creating new ones.
	// Temporarily hardcoding the collection, this needs to be changed to reflect the cluster it belongs to later.
	if _, err := client.Collection("prod").Doc(server.ID).Set(ctx, server); err != nil {
		return types.Server{}, fmt.Errorf("error writing to Firestore database:\n%v", err)
	}

	return server, nil
}

package db

import (
	"context"
	"fmt"
	"log"

	"github.com/RicochetStudios/aurora/types"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

const (
	// instancePath is the path to the instance documents.
	instancePath string = "default/instances/"

	// dbUrl is the url for the Firebase database instance.
	dbUrl string = "https://game-server-e2c56-default-rtdb.europe-west1.firebasedatabase.app"

	// dbAuth is the authentication config path, used to access firebase.
	dbAuth string = "./firebase-config.json"
)

// Call this function to get/initialise the firebase app
func Firebase(ctx context.Context) (*firebase.App, error) {
	// configure database URL
	conf := &firebase.Config{
		DatabaseURL: dbUrl,
	}

	// fetch service account key
	opt := option.WithCredentialsFile(dbAuth)

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	return app, err
}

// Call this function to get a client for the realtime database
func RealtimeDatabase(ctx context.Context) (*db.Client, error) {
	app, err := Firebase(ctx)

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
func Firestore(ctx context.Context) (*firestore.Client, error) {
	app, err := Firebase(ctx)

	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln("error in creating firebase DB client: ", err)
	}

	return client, err
}

// GetServer reads and returns a server document, given an ID.
func GetServer(ctx context.Context, id string) (types.Server, error) {
	// Create the firestore client.
	client, err := Firestore(ctx)
	if err != nil {
		return types.Server{}, fmt.Errorf("error creating Firestore client:\n%v", err)
	}
	defer client.Close()

	// Read the full document from the database.
	// Temporarily hardcoding the collection, this needs to be changed to reflect the cluster it belongs to later.
	document, err := client.Collection("development").Doc(instancePath + id).Get(ctx)
	if err != nil {
		return types.Server{}, fmt.Errorf("error reading document from Firestore database:\n%v", err)
	}

	// Convert the document into the server struct.
	var server types.Server
	if err := document.DataTo(&server); err != nil {
		return types.Server{}, fmt.Errorf("error converting Firestore document to types.Server struct:\n%v", err)
	}

	return server, nil
}

// SetServer creates and overwrites fields in the server document, given a Server.
func SetServer(ctx context.Context, id string, server types.Server) (types.Server, error) {
	// Create the firestore client.
	client, err := Firestore(ctx)
	if err != nil {
		return types.Server{}, fmt.Errorf("error creating Firestore client:\n%v", err)
	}
	defer client.Close()

	// Write to the database, overwriting existing fields and creating new ones.
	// Temporarily hardcoding the collection, this needs to be changed to reflect the cluster it belongs to later.
	if _, err := client.Collection("development").Doc(instancePath+id).Set(ctx, server); err != nil {
		return types.Server{}, fmt.Errorf("error writing to document in Firestore database:\n%v", err)
	}

	return server, nil
}

// RemoveServer removes an instance document from the database, given an ID.
func RemoveServer(ctx context.Context, id string) error {
	// Create the firestore client.
	client, err := Firestore(ctx)
	if err != nil {
		return fmt.Errorf("error creating Firestore client:\n%v", err)
	}
	defer client.Close()

	// Removing the server instance from the database by deleting the corresponding document.
	// Temporarily hardcoding the collection, this needs to be changed to reflect the cluster it belongs to later.
	if _, err := client.Collection("development").Doc(instancePath + id).Delete(ctx); err != nil {
		return fmt.Errorf("error deleting document from Firestore database:\n%v", err)
	}

	return nil
}

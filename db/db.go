package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

// Call this function to get/initialise the firebase app
func FireBase() (*firebase.App, error) {
	ctx := context.Background()

	// configure database URL
	conf := &firebase.Config{
		DatabaseURL: "https://game-server-e2c56-default-rtdb.europe-west1.firebasedatabase.app",
	}

	// fetch service account key
	opt := option.WithCredentialsFile("db/config/game-server-e2c56-firebase-adminsdk-thy9x-ab047e2cfa.json")


	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	return app, err
}

// Call this function to get a client for the realtime database
func RealtimeDatabase() (*db.Client, error) {
	ctx := context.Background()

	app, err := FireBase();

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

	app, err := FireBase();

	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln("error in creating firebase DB client: ", err)
	}

	return client, err
}



	//print type of client
	// return client

	// // create ref at path user_scores/:userId
	// ref := client.NewRef("user_scores/" + fmt.Sprint(1))

	// if err := ref.Set(context.TODO(), map[string]interface{}{"score": 40}); err != nil {
	// log.Fatal(err)
	// }

	// fmt.Println("score added/updated successfully!")
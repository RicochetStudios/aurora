package db

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

// have client in return type
func FireBaseInstance() (*db.Client, error) {
	ctx := context.Background()

	// configure database URL
	conf := &firebase.Config{
		DatabaseURL: "https://game-server-e2c56-default-rtdb.europe-west1.firebasedatabase.app",
	}

	// fetch service account key
	opt := option.WithCredentialsFile("db/game-server-e2c56-firebase-adminsdk-thy9x-ab047e2cfa.json")


	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	client, err := app.Database(ctx)
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
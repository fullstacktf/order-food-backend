package database

import (
	"comiditapp/api/env"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createClient() (*mongo.Client, *mongo.Database) {
	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	if err != nil {
		log.Fatal("Error creating mongo client ❌: ", err)
	}

	println("Mongo client created ✅")

	if err = client.Connect(context.TODO()); err != nil {
		log.Fatal("Error when connecting to database ❌: ", err)
	}

	println("Connected to the database ✅")

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal("Error when trying to ping database ❌: ", err)
	}

	println("Ping done succesfully ✅")

	var db = client.Database(env.DB_NAME)

	return client, db
}

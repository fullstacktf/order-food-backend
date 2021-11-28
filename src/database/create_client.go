package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createClient() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	if err != nil {
		log.Fatal("Error when connecting to database ❌: ", err)
	}

	println("Connected to the database ✅")

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal("Error when trying to ping database ❌: ", err)
	}

	println("Ping done succesfully ✅")
	return client
}

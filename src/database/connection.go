package database

import (
	"comiditapp/api/env"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection_names []string = []string{"user", "order"}

type DB struct {
	Client      *mongo.Client
	Collections map[string]*mongo.Collection
}

func (db *DB) Init() {
	db.Client = createClient()
	println("Client initialized succesfully ✅")

	db.Collections = getAllCollections(db.Client)
	println("Collections initialized succesfully ✅")
}

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

func getAllCollections(client *mongo.Client) map[string]*mongo.Collection {

	collections := make(map[string]*mongo.Collection)

	for _, collectionName := range collection_names {
		collections[collectionName] = getCollection(client, collectionName)
	}

	return collections
}

func getCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(env.DB_NAME).Collection(collectionName)
}

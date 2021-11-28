package database

import (
	"comiditapp/api/env"

	"go.mongodb.org/mongo-driver/mongo"
)

var collection_names []string = []string{"user", "order"}

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

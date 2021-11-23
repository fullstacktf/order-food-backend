package database

import (
	"comiditapp/api/src/env"
	"context"
	"log"
	"os"
	"time"

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
	db.Collections = getAllCollections(db.Client)
}

func (db *DB) DropDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db.Client.Database(env.DB_NAME).Drop(ctx)
}

func createClient() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
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

func GetDB() DB {
	var db DB
	db.Init()
	return db
}

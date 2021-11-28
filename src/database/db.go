package database

import "go.mongodb.org/mongo-driver/mongo"

type DB struct {
	Client      *mongo.Client
	Collections map[string]*mongo.Collection
}

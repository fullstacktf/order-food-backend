package database

import (
	orders_repository "comiditapp/api/database/repositories/orders"
	users_repository "comiditapp/api/database/repositories/users"

	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	Client           *mongo.Client
	Database         *mongo.Database
	Collections      map[string]*mongo.Collection
	OrdersRepository orders_repository.MongoOrdersRepository
	UsersRepository  users_repository.MongoUsersRepository
}

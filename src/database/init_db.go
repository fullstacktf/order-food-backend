package database

import (
	orders_repository "comiditapp/api/database/repositories/orders"
	users_repository "comiditapp/api/database/repositories/users"
)

func (db *DB) Init() {
	db.Client, db.Database = createClient()
	println("Client and database initialized succesfully ✅")

	ordersRepo := orders_repository.NewMongoOrdersRepository(db.Database)
	usersRepo := users_repository.NewMongoUsersRepository(db.Database)

	db.OrdersRepository, db.UsersRepository = *ordersRepo, *usersRepo

	db.Collections = getAllCollections(db.Client)
	println("Collections initialized succesfully ✅")
}

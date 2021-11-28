package database

func (db *DB) Init() {
	db.Client = createClient()
	println("Client initialized succesfully ✅")

	db.Collections = getAllCollections(db.Client)
	println("Collections initialized succesfully ✅")
}

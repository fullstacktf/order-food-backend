package database

func GetDB() DB {
	var db DB

	db.Init()

	return db
}

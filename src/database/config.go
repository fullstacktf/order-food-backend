package database

import(
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const MONGODB_URI = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000"
const DB_NAME = "comiditapp_db"

func init() {
	// Setup de mgm default config
	err := mgm.SetDefaultConfig(nil, DB_NAME, options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		log.Fatal(err)
	}
}


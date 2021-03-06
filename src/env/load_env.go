package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var MONGODB_URI string
var DB_NAME string
var PORT string
var SECRET string

func LoadEnv() {

	if os.Getenv("ENV") != "PROD" {
		err := godotenv.Load("./.env")

		if err != nil {
			log.Fatal("Error loading .env file ❌: ", err)
		}
	}

	MONGODB_URI = os.Getenv("MONGODB_URI")
	DB_NAME = os.Getenv("DB_NAME")
	PORT = os.Getenv("PORT")
	SECRET = os.Getenv("SECRET")

	println("Loaded environment variables succesfully ✅")
}

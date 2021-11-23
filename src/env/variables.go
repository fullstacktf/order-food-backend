package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var MONGODB_URI string
var DB_NAME string
var PORT string

func LoadEnv() {

	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	MONGODB_URI = os.Getenv("MONGODB_URI")
	DB_NAME = os.Getenv("DB_NAME")
	PORT = os.Getenv("PORT")
}

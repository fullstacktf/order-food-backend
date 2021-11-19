package main

import (
	auth "comiditapp/api/src/controllers/auth"
	client "comiditapp/api/src/controllers/client"
	profile "comiditapp/api/src/controllers/profile"
	restaurant "comiditapp/api/src/controllers/restaurant"
	database "comiditapp/api/src/database"
	"comiditapp/api/src/env"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	var db database.DB
	db.Init()

	r := gin.Default()

	auth.Routes(r)
	profile.Routes(r)
	client.Routes(r)
	restaurant.Routes(r)

	r.Run()

	// Server listening on port 3000
	if err := r.Run(env.PORT); err != nil {
		log.Fatal(err.Error())
	}
}

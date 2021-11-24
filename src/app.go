package main

import (
	"comiditapp/api/controllers/auth"
	"comiditapp/api/controllers/client"
	"comiditapp/api/controllers/profile"
	"comiditapp/api/controllers/restaurant"
	"comiditapp/api/database"
	"comiditapp/api/env"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	env.LoadEnv()

	var db database.DB
	db.Init()

	r := gin.Default()

	auth.Routes(r)
	profile.Routes(r)
	client.Routes(r)
	restaurant.Routes(r)

	r.Run()

	// Server listening on port 3000
	if err := r.Run(":" + env.PORT); err != nil {
		log.Fatal("Error running the server ❌: ", err.Error())
	}
	println("yeka❌❌❌❌❌❌❌❌")
}

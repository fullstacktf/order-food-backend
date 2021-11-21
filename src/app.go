package main

import (
	"comiditapp/api/src/controllers/auth"
	"comiditapp/api/src/controllers/client"
	"comiditapp/api/src/controllers/profile"
	"comiditapp/api/src/controllers/restaurant"
	"comiditapp/api/src/env"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	//db := database.GetDB()

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

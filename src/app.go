package main

import (
	"log"

	auth "comiditapp/api/src/controllers/auth"
	client "comiditapp/api/src/controllers/client"
	profile "comiditapp/api/src/controllers/profile"
	restaurant "comiditapp/api/src/controllers/restaurant"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	auth.Routes(r)
	profile.Routes(r)
	client.Routes(r)
	restaurant.Routes(r)

	r.Run()

	// Server listening on port 3000
	if err := r.Run(":3000"); err != nil {
		log.Fatal(err.Error())
	}
}

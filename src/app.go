package main

import (
	"log"

	auth "comiditapp/api/src/routes/auth"
	client "comiditapp/api/src/routes/client"
	profile "comiditapp/api/src/routes/profile"
	restaurant "comiditapp/api/src/routes/restaurant"

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

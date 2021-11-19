package main

import (
	"log"
	"os"
  "github.com/gin-gonic/gin"
	auth "comiditapp/api/src/controllers/auth"
	client "comiditapp/api/src/controllers/client"
	profile "comiditapp/api/src/controllers/profile"
	restaurant "comiditapp/api/src/controllers/restaurant"
)

func main() {
	r := gin.Default()

	auth.Routes(r)
	profile.Routes(r)
	client.Routes(r)
	restaurant.Routes(r)

	r.Run()

	// Server listening on port 3000
	if err := r.Run(os.Getenv("PORT")); err != nil {
		log.Fatal(err.Error())
	}
}


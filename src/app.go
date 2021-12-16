package main

import (
	auth "comiditapp/api/controllers/auth"
	clients "comiditapp/api/controllers/clients"
	home "comiditapp/api/controllers/home"
	orders "comiditapp/api/controllers/orders"
	products "comiditapp/api/controllers/products"
	profiles "comiditapp/api/controllers/profile"
	restaurants "comiditapp/api/controllers/restaurants"
	"comiditapp/api/database"
	"comiditapp/api/env"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	env.LoadEnv()

	db := database.GetDB()
	db.Init()
	//database.SetInitialData(db)

	r := gin.Default()
	r.Use(cors.Default())

	home.Routes(r)
	auth.Routes(r, db)
	restaurants.Routes(r, db)
	clients.Routes(r, db)
	profiles.Routes(r, db)
	orders.Routes(r, db)
	products.Routes(r, db)

	r.Run()

	// Server listening on port 3000
	if err := r.Run(env.PORT); err != nil {
		log.Fatal("Error running the server ❌: ", err.Error())
	}
}

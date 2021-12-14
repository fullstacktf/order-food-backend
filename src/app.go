package main

import (
	"comiditapp/api/controllers/any_role/auth"
	client "comiditapp/api/controllers/any_role/clients"
	"comiditapp/api/controllers/any_role/home"
	"comiditapp/api/controllers/any_role/profile"
	"comiditapp/api/controllers/any_role/restaurants"
	"comiditapp/api/controllers/restaurant_role/orders"
	"comiditapp/api/controllers/restaurant_role/products"
	"comiditapp/api/database"
	"comiditapp/api/env"
	"log"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	env.LoadEnv()

	db := database.GetDB()
	db.Init()
	//database.SetInitialData(db)

	r := gin.Default()
	r.Use(cors.Default())

	// any role users
	home.Routes(r)
	auth.Routes(r, db)
	restaurants.Routes(r, db)
	client.Routes(r, db)
	profile.Routes(r, db)

	// only restaurant role users
	orders.Routes(r, db)
	products.Routes(r, db)

	r.Run()

	// Server listening on port 3000
	if err := r.Run(env.PORT); err != nil {
		log.Fatal("Error running the server ❌: ", err.Error())
	}
}

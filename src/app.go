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
	"comiditapp/api/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	env.LoadEnv()

	db := database.GetDB()
	db.Init()
	//database.SetInitialData(db)

	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:8080"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "http://localhost:8080"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))

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

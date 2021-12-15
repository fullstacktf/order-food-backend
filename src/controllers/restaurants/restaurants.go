package controllers

import (
	"comiditapp/api/database"
	handlers "comiditapp/api/handlers/restaurants"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db database.DB) {
	restaurantGroup := r.Group("/restaurants")
	{
		restaurantGroup.POST("/:id/orders", handlers.CreateOrder(db.OrdersRepository))
		restaurantGroup.GET("", handlers.FindRestaurants(db.UsersRepository))
		restaurantGroup.GET("/:id", handlers.GetRestaurantById(db.UsersRepository))
		restaurantGroup.GET("/:id/products", handlers.GetRestaurantProducts(db.UsersRepository))
	}
}

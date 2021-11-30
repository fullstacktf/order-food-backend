package restaurants

import (
	"comiditapp/api/database"
	restaurant_handler "comiditapp/api/handlers/any_role/restaurants"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db database.DB) {
	restaurantGroup := r.Group("/restaurants")
	{
		restaurantGroup.POST("/:id/orders", restaurant_handler.CreateOrder(db.OrdersRepository))

		restaurantGroup.GET("", restaurant_handler.FindRestaurants(db.UsersRepository))
		restaurantGroup.GET("/:id", restaurant_handler.GetRestaurantById(db.UsersRepository))
		restaurantGroup.GET("/:id/products", restaurant_handler.GetRestaurantProducts(db.UsersRepository))
	}
}

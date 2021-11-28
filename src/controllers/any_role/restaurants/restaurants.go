package restaurants

import (
	orders_repository "comiditapp/api/database/repositories/orders"
	users_repository "comiditapp/api/database/repositories/users"
	restaurant_handler "comiditapp/api/handlers/any_role/restaurants"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	var ordersRepository orders_repository.MockedOrdersRepository
	var usersRepository users_repository.MockedUsersRepository

	restaurantGroup := r.Group("/restaurants")
	{
		restaurantGroup.POST("/:id/orders", restaurant_handler.CreateOrder(ordersRepository))

		restaurantGroup.GET("", restaurant_handler.FindRestaurants(usersRepository))
		restaurantGroup.GET("/:id", restaurant_handler.GetRestaurantById(usersRepository))
		restaurantGroup.GET("/:id/products", restaurant_handler.GetRestaurantProducts(usersRepository))
	}
}

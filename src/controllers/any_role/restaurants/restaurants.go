package restaurants

import (
	repository "comiditapp/api/database/repositories/orders"
	restaurant_handler "comiditapp/api/handlers/any_role/restaurants"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	var ordersRepository repository.MockedOrdersRepository

	restaurantGroup := r.Group("/restaurants")
	{
		restaurantGroup.POST("/:id/orders", restaurant_handler.CreateOrder(ordersRepository))
		// restaurantGroup.GET("", restaurant_handler.GetAllRestaurants)
		// restaurantGroup.GET("/:id", restaurant_handler.GetRestaurantById)
		// restaurantGroup.GET("/:id/products", restaurant_handler.GetRestaurantProducts)
	}
}

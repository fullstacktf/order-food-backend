package restaurants

import (
	restaurant_handler "comiditapp/api/handlers/any_role/restaurants"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	restaurantGroup := r.Group("/restaurants")
	{
		restaurantGroup.GET("", restaurant_handler.GetAllRestaurants)
		restaurantGroup.GET("/:id", restaurant_handler.GetRestaurantById)
		restaurantGroup.GET("/:id/products", restaurant_handler.GetRestaurantProducts)
		restaurantGroup.POST("/:id/orders", restaurant_handler.PostOrder)
	}
}

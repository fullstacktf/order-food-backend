package restaurant

import (
	handler "comiditapp/api/handlers/restaurant"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	restaurant := r.Group("restaurant")
	restaurant.GET("/products", handler.GetRestaurantProductsForRestaurant)
	restaurant.POST("/products", handler.PostRestaurantProduct)
	restaurant.PUT("/products/:id", handler.UpdateRestaurantProductById)
}

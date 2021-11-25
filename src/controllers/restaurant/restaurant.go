package restaurant

import (
	handler "comiditapp/api/handlers/restaurant"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	// restaurant := r.Group("restaurant")
	r.GET("/products", handler.GetRestaurantProductsForRestaurant)
	r.POST("/products", handler.PostRestaurantProduct)
	r.PUT("/products/:id", handler.UpdateRestaurantProductById)
}

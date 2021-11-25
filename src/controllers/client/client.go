package client

import (
	handler "comiditapp/api/handlers/client"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	// client := r.Group("client")
	r.GET("/", handler.GetHomePage)
	r.GET("/restaurants", handler.GetAllRestaurants)
	r.GET("/restaurants/:id", handler.GetRestaurantById)
	r.GET("/restaurants/:id/products", handler.GetRestaurantProductsForClient)
	r.POST("/order", handler.PostClientOrder)
}

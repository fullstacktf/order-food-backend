package client

import (
	handler "comiditapp/api/handlers/client"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	client := r.Group("client")
	client.GET("/", handler.GetHomePage)
	client.GET("/home", handler.GetHomePage)
	client.GET("/restaurants", handler.GetAllRestaurants)
	client.GET("/restaurants/:id", handler.GetRestaurantById)
	client.GET("/restaurants/:id/products", handler.GetRestaurantProductsForClient)
	client.POST("/order", handler.PostClientOrder)
}

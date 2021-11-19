package restaurant

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	restaurant := r.Group("restaurant")
	restaurant.GET("/products", GetRestaurantProductsForRestaurant)
	restaurant.POST("/products", PostRestaurantProduct)
	restaurant.PUT("/products/:id", UpdateRestaurantProductById)
}

func GetRestaurantProductsForRestaurant(c *gin.Context) {} // No tenemos id como parametro, buscar manera de asociar al id de restaurante
func PostRestaurantProduct(c *gin.Context)              {} // No tenemos id como parametro, buscar manera de asociar al id de restaurante
func UpdateRestaurantProductById(c *gin.Context)        {} // No tenemos id como parametro, buscar manera de asociar al id de restaurante

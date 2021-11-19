package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	client := r.Group("client")
	client.GET("/", GetHomePage)
	client.GET("/home", GetHomePage)
	client.GET("/restaurants", GetAllRestaurants)
	client.GET("/restaurants/:id", GetRestaurantById)
	client.GET("/restaurants/:id/products", GetRestaurantProductsForClient)
	client.POST("/order", PostClientOrder)
}

func GetHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "You are at comiditapp homepage🥳🥳",
	})
}

func GetAllRestaurants(c *gin.Context) {
	var restaurants []string
	c.JSON(200, gin.H{
		"restaurants": restaurants,
	})
}

func GetRestaurantById(c *gin.Context) {
	name := c.Param("id")
	// Logica para buscar el restaurante en la bbdd y devolver sus datos
	// con el modelo de respuesta que definamos
	// ...
	message := "restaurant" + name + "was found"
	c.String(http.StatusOK, message)
}

func GetRestaurantProductsForClient(c *gin.Context) {
	name := c.Param("id")
	// Logica para traer todos los productos de un restaurante de la bbdd
	// y devolverlos con el modelo definido para ellos
	// ...
	message := "restaurant" + name + "was found"
	c.String(http.StatusOK, message)
}

func PostClientOrder(c *gin.Context) {}

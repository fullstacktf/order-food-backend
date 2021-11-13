package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	// All the controllers

	// Clients
	r.GET("/", GetHomePage)
	r.GET("/home", GetHomePage)
	r.GET("/restaurants", GetAllRestaurants)
	r.GET("/restaurants/:id", GetRestaurantById)
	r.GET("/restaurants/:id/products", GetRestaurantProducts)
	r.POST("/auth/signup/client", PostSignUpClient)
	r.POST("/auth/signin/client", PostSignInClient)
	r.POST("/order", PostClientOrder)

	// Both
	r.GET("/profile/orders", GetProfileOrders)
	r.GET("/profile/orders/:id", GetProfileOrderById)
	r.PUT("/profile", PutProfileData)

	//  Restaurants
	r.GET("/products", PostSignUpRestaurant)
	r.POST("/auth/signup/restaurant", PostSignUpRestaurant)
	r.POST("/auth/signin/restaurant", PostSignInRestaurant)
	r.POST("/products", PostSignUpRestaurant)
	r.PUT("/products/:id", PutRestaurantProductData)

	// Server listening on port 3000
	if err := r.Run(":3000"); err != nil {
		log.Fatal(err.Error())
	}
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

func GetRestaurantProducts(c *gin.Context) {
	name := c.Param("id")
	// Logica para traer todos los productos de un restaurante de la bbdd
	// y devolverlos con el modelo definido para ellos
	// ...
	message := "restaurant" + name + "was found"
	c.String(http.StatusOK, message)
}

func GetProfileOrders(c *gin.Context) {
	name := c.Param("id")
	// Logica
	message := "all good" + name
	c.String(http.StatusOK, message)
}

func GetProfileOrderById(c *gin.Context) {
	name := c.Param("id")
	// Logica
	message := "all good" + name
	c.String(http.StatusOK, message)
}

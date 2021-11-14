package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// All the controllers

	// Clients
	r.GET("/", GetHomePage)
	r.GET("/home", GetHomePage)
	r.GET("/restaurants", GetAllRestaurants)
	r.GET("/restaurants/:id", GetRestaurantById)
	r.GET("/restaurants/:id/products", GetRestaurantProductsForClient)
	r.POST("/auth/signup/client", PostSignUpClient)
	r.POST("/auth/signin/client", PostSignInClient)
	r.POST("/order", PostClientOrder)

	// Both
	r.GET("/profile/orders", GetProfileOrders)
	r.GET("/profile/orders/:id", GetProfileOrderById)
	r.PUT("/profile", UpdateProfileData)

	//  Restaurants
	r.POST("/auth/signup/restaurant", PostSignUpRestaurant)
	r.POST("/auth/signin/restaurant", PostSignInRestaurant)
	r.GET("/products", GetRestaurantProductsForRestaurant)
	r.POST("/products", PostRestaurantProduct)
	r.PUT("/products/:id", UpdateRestaurantProductById)

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

func GetRestaurantProductsForClient(c *gin.Context) {
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

func PostSignUpClient(c *gin.Context)                   {}
func PostSignInClient(c *gin.Context)                   {}
func PostClientOrder(c *gin.Context)                    {}
func UpdateProfileData(c *gin.Context)                  {}
func PostSignUpRestaurant(c *gin.Context)               {}
func PostSignInRestaurant(c *gin.Context)               {}
func GetRestaurantProductsForRestaurant(c *gin.Context) {} // No tenemos id como parametro, buscar manera de asociar al id de restaurante
func PostRestaurantProduct(c *gin.Context)              {} // No tenemos id como parametro, buscar manera de asociar al id de restaurante
func UpdateRestaurantProductById(c *gin.Context)        {} // No tenemos id como parametro, buscar manera de asociar al id de restaurante

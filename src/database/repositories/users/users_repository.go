package repository

import "github.com/gin-gonic/gin"

type UsersRepository interface {
	// any_role methods
	SignUpUser(context *gin.Context)            // Endpoint -> /auth/signup
	SignInUser(context *gin.Context)            // Endpoint -> /auth/signin
	FindRestaurants(context *gin.Context)       // Endpoint -> /restaurants
	GetRestaurantById(context *gin.Context)     // Endpoint -> /restaurants/:id
	GetRestaurantProducts(context *gin.Context) // Endpoint -> /restaurants/:id/products
	UpdateProfile(context *gin.Context)         // Endpoint -> /profile

	// restaurant_role methods
	FindProducts(context *gin.Context)  // Endpoint -> /products
	CreateProduct(context *gin.Context) // Endpoint -> /products
	UpdateProduct(context *gin.Context) // Endpoint -> /products/:id
}

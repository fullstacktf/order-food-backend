package products

import (
	repository "comiditapp/api/database/repositories/users"
	product_handler "comiditapp/api/handlers/restaurant_role/products"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	var usersRepository repository.MockedUsersRepository

	productsGroup := r.Group("/products")
	{
		productsGroup.GET("", product_handler.FindProducts(usersRepository))
		productsGroup.POST("", product_handler.CreateProduct(usersRepository))
		productsGroup.PUT("/:id", product_handler.UpdateProduct(usersRepository))
	}
}

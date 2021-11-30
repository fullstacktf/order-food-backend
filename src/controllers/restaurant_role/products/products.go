package products

import (
	"comiditapp/api/database"
	product_handler "comiditapp/api/handlers/restaurant_role/products"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db database.DB) {
	productsGroup := r.Group("/products")
	{
		productsGroup.GET("", product_handler.FindProducts(db.UsersRepository))
		productsGroup.POST("", product_handler.CreateProduct(db.UsersRepository))
		productsGroup.PUT("/:id", product_handler.UpdateProduct(db.UsersRepository))
	}
}

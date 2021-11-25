package products

import (
	product_handler "comiditapp/api/handlers/restaurant_role/products"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	productsGroup := r.Group("/products")
	{
		productsGroup.GET("", product_handler.GetAllProducts)
		productsGroup.POST("", product_handler.PostProduct)
		productsGroup.PUT("/:id", product_handler.UpdateProduct)
	}
}

package controllers

import (
	"comiditapp/api/database"
	handlers "comiditapp/api/handlers/products"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db database.DB) {
	productsGroup := r.Group("/products")
	{
		productsGroup.GET("", handlers.FindProducts(db.UsersRepository))
		productsGroup.POST("", handlers.CreateProduct(db.UsersRepository))
		productsGroup.PUT("/:id", handlers.UpdateProduct(db.UsersRepository))
	}
}

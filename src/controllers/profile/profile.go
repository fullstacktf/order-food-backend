package controllers

import (
	"comiditapp/api/database"
	handlers "comiditapp/api/handlers/profile"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db database.DB) {
	profileGroup := r.Group("/profile")
	{
		profileGroup.PUT("/:id", handlers.UpdateProfile(db.UsersRepository))
		profileGroup.GET("/:id", handlers.GetProfileById(db.UsersRepository))
		profileGroup.DELETE("/:id", handlers.DeleteAccount(db.UsersRepository))
		profileGroup.GET("/orders", handlers.FindOrders(db.OrdersRepository))
		profileGroup.GET("/orders/:id", handlers.GetOrderById(db.OrdersRepository))
	}
}

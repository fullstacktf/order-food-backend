package controllers

import (
	"comiditapp/api/database"
	handlers "comiditapp/api/handlers/orders"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db database.DB) {
	ordersGroup := r.Group("/orders")
	{
		ordersGroup.GET("", handlers.FindClientOrders(db.OrdersRepository))
		ordersGroup.PUT("/:id", handlers.UpdateClientOrder(db.OrdersRepository))
	}
}

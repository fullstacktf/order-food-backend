package orders

import (
	"comiditapp/api/database"
	order_handler "comiditapp/api/handlers/restaurant_role/orders"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db database.DB) {
	ordersGroup := r.Group("/orders")
	{
		ordersGroup.GET("", order_handler.FindClientOrders(db.OrdersRepository))
		ordersGroup.PUT("/:id", order_handler.UpdateClientOrder(db.OrdersRepository))
	}
}

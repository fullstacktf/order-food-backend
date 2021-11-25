package orders

import (
	order_handler "comiditapp/api/handlers/restaurant_role/orders"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	ordersGroup := r.Group("/orders")
	{
		ordersGroup.GET("", order_handler.GetAllOrders)
		ordersGroup.PUT("/:id", order_handler.UpdateOrder)
	}
}

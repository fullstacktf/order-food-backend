package orders

import (
	repository "comiditapp/api/database/repositories/orders"
	order_handler "comiditapp/api/handlers/restaurant_role/orders"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	var ordersRepository repository.MockedOrdersRepository

	ordersGroup := r.Group("/orders")
	{
		ordersGroup.GET("", order_handler.GetClientsOrders(ordersRepository))
		ordersGroup.PUT("/:id", order_handler.UpdateClientOrder(ordersRepository))
	}
}

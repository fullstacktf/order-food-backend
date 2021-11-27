package profile

import (
	repository "comiditapp/api/database/repositories/orders"
	profile_handler "comiditapp/api/handlers/any_role/profile"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	var ordersRepository repository.MockedOrdersRepository

	profileGroup := r.Group("/profile")
	{
		profileGroup.PUT("", profile_handler.UpdateProfile)
		profileGroup.GET("/orders", profile_handler.GetOrders(ordersRepository))
		profileGroup.GET("/orders/:id", profile_handler.GetOrderById(ordersRepository))
	}
}

package profile

import (
	orders_repository "comiditapp/api/database/repositories/orders"
	users_repository "comiditapp/api/database/repositories/users"

	profile_handler "comiditapp/api/handlers/any_role/profile"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	var ordersRepository orders_repository.MockedOrdersRepository
	var usersRepository users_repository.MockedUsersRepository

	profileGroup := r.Group("/profile")
	{
		// This endpoint is temporary using id as we still working in register and login
		profileGroup.PUT("/:id", profile_handler.UpdateProfile(usersRepository))

		profileGroup.GET("/orders", profile_handler.GetOrders(ordersRepository))
		profileGroup.GET("/orders/:id", profile_handler.GetOrderById(ordersRepository))
	}
}

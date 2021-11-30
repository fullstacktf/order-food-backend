package profile

import (
	"comiditapp/api/database"

	profile_handler "comiditapp/api/handlers/any_role/profile"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db database.DB) {
	profileGroup := r.Group("/profile")
	{
		// This endpoint is temporary using id as we still working in register and login
		profileGroup.PUT("/:id", profile_handler.UpdateProfile(db.UsersRepository))

		profileGroup.GET("/orders", profile_handler.FindOrders(db.OrdersRepository))
		profileGroup.GET("/orders/:id", profile_handler.GetOrderById(db.OrdersRepository))
	}
}

package profile

import (
	profile_handler "comiditapp/api/handlers/any_role/profile"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	profileGroup := r.Group("/profile")
	{
		profileGroup.PUT("", profile_handler.UpdateProfile)
		profileGroup.GET("/orders", profile_handler.GetOrders)
		profileGroup.GET("/orders/:id", profile_handler.GetOrderById)
	}
}

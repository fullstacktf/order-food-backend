package profile

import (
	handler "comiditapp/api/handlers/profile"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	// profile := r.Group("profile")
	r.GET("/profile/orders", handler.GetProfileOrders)
	r.GET("/profile/orders/:id", handler.GetProfileOrderById)
	r.PUT("/profile", handler.UpdateProfileData)
}

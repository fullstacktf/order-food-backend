package profile

import (
	handler "comiditapp/api/src/handlers/profile"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	profile := r.Group("profile")
	profile.GET("/profile/orders", handler.GetProfileOrders)
	profile.GET("/profile/orders/:id", handler.GetProfileOrderById)
	profile.PUT("/profile", handler.UpdateProfileData)
}

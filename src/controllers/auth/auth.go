package auth

import (
	handler "comiditapp/api/src/handlers/auth"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	auth := r.Group("auth")
	auth.POST("/auth/signup/restaurant", handler.PostSignUpRestaurant)
	auth.POST("/auth/signin/restaurant", handler.PostSignInRestaurant)
	auth.POST("/auth/signup/client", handler.PostSignUpClient)
	auth.POST("/auth/signin/client", handler.PostSignInClient)
}

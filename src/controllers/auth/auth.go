package auth

import (
	handler "comiditapp/api/handlers/auth"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	// auth := r.Group("auth")
	r.POST("/auth/signup/restaurant", handler.PostSignUpRestaurant)
	r.POST("/auth/signin/restaurant", handler.PostSignInRestaurant)
	r.POST("/auth/signup/client", handler.PostSignUpClient)
	r.POST("/auth/signin/client", handler.PostSignInClient)
}

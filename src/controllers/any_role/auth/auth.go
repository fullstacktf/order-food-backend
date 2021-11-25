package auth

import (
	auth_handler "comiditapp/api/handlers/any_role/auth"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", auth_handler.PostSignUp)
		authGroup.POST("/signin", auth_handler.PostSignIn)
	}
}

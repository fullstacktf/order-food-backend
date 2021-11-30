package auth

import (
	"comiditapp/api/database"
	auth_handler "comiditapp/api/handlers/any_role/auth"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db database.DB) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", auth_handler.SignUpUser(db.UsersRepository))
		authGroup.POST("/signin", auth_handler.SignInUser(db.UsersRepository))
	}
}

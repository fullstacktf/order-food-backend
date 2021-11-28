package auth

import (
	repository "comiditapp/api/database/repositories/users"
	auth_handler "comiditapp/api/handlers/any_role/auth"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	var usersRepository repository.MockedUsersRepository

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", auth_handler.SignUpUser(usersRepository))
		authGroup.POST("/signin", auth_handler.SignInUser(usersRepository))
	}
}

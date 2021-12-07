package auth_handler

import (
	repository "comiditapp/api/database/repositories/users"

	"github.com/gin-gonic/gin"
)

func SignInUser(repository *repository.MongoUsersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		statusCode, response := repository.SignInUser(context)
		context.IndentedJSON(statusCode, response)

	}
}

package client_handler

import (
	repository "comiditapp/api/database/repositories/users"

	"github.com/gin-gonic/gin"
)

func FindClients(repository *repository.MongoUsersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		statusCode, response := repository.FindClients(context)
		context.IndentedJSON(statusCode, response)
	}
}

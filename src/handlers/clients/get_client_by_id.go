package handlers

import (
	repository "comiditapp/api/database/repositories/users"

	"github.com/gin-gonic/gin"
)

func GetClientById(repository *repository.MongoUsersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		statusCode, response := repository.GetClientById(context)
		context.IndentedJSON(statusCode, response)
	}
}
